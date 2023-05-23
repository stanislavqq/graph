package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"graph/config"
	"net/http"
	"time"
)

type Server struct {
	Router     *mux.Router
	Logger     zerolog.Logger
	HttpConfig config.HttpConfig
}

func New(cfg config.HttpConfig, logger zerolog.Logger) *Server {
	srv := &Server{
		Router:     mux.NewRouter(),
		Logger:     logger,
		HttpConfig: cfg,
	}

	//srv.Router.HandleFunc("/getChart", chartHandler).Methods("GET")
	//srv.Router.HandleFunc("/", homeHandler).Methods("GET")

	return srv
}

func (s *Server) Start() error {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("%s:%s", s.HttpConfig.Host, s.HttpConfig.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		s.Logger.Fatal().Err(err)
		return err
	}

	return nil
}
