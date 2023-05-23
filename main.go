package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"graph/app"
	"graph/config"
	"graph/parser"
	"graph/server"
	"net/http"
	"os"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.Logger{}
	err := godotenv.Load()
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading .env file")
	}
}

func main() {
	CSVParser, err := parser.NewParser("csv")
	CSVParser.SetSkipRows(2)
	if err != nil {
		logger.Fatal().Err(err)
		os.Exit(1)
	}

	MChart := app.NewChart(CSVParser, logger)
	appChart := app.NewApp(MChart)

	cfg := config.Config{}
	cfg.HttpConfig = config.HttpConfig{
		Host: "127.0.0.1",
		Port: "8000",
	}

	srv := server.New(cfg.HttpConfig, logger)

	srv.Router.HandleFunc("/getChart", appChart.GetChart).Methods("GET")
	srv.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	}).Methods("GET")

	if err = srv.Start(); err != nil {
		logger.Fatal().Err(err)
		return
	}
}
