package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"graph/app"
	"graph/config"
	"graph/parser"
	"graph/server"
	"net/http"
	"os"
	"os/signal"
)

var logger zerolog.Logger

func init() {
	logger = log.Logger
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
	cfg.HttpConfig = config.HttpConfig{Host: "", Port: ""}

	host := flag.String("host", "127.0.0.1", "-host 127.0.0.1")
	port := flag.String("port", "8000", "-port 8000")
	flag.Parse()

	cfg.HttpConfig.Host = *host
	cfg.HttpConfig.Port = *port

	fmt.Println(cfg.HttpConfig.Host)

	ctx := context.Background()
	srv := server.New(cfg.HttpConfig, logger)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	srv.Router.HandleFunc("/getChart", appChart.GetChart).Methods("GET")
	srv.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	}).Methods("GET")

	go func() {
		select {
		case <-done:
			logger.Info().Msg("Get exit signal")
			logger.Info().Msg("Http server close")
			defer srv.Stop()
			//os.Exit(1)
		}
	}()

	if err = srv.Start(ctx); err != nil {
		logger.Fatal().Err(err)
		return
	}
}
