package server

import (
	"context"
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
	HttpServer *http.Server
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

func (s *Server) Start(ctx context.Context) error {
	s.HttpServer = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("%s:%s", s.HttpConfig.Host, s.HttpConfig.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.Logger.Info().Msg(fmt.Sprintf("Http Server started on %s:%s", s.HttpConfig.Host, s.HttpConfig.Port))

	if err := s.HttpServer.ListenAndServe(); err != nil {
		s.Logger.Fatal().Err(err)
		return err
	}

	return nil
}

func (s *Server) Stop() {
	if err := s.HttpServer.Close(); err != nil {
		s.Logger.Error().Err(err)
		return
	}
}
