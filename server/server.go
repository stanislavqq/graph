package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"graph/app"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Router *mux.Router
}

func New() *Server {
	srv := &Server{
		Router: mux.NewRouter(),
	}

	srv.Router.HandleFunc("/getChart", chartHandler).Methods("GET")
	srv.Router.HandleFunc("/", homeHandler).Methods("GET")

	return srv
}

func (s *Server) Start() {

	//http.Handle("/", router)

	srv := &http.Server{
		Handler: s.Router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func chartHandler(w http.ResponseWriter, r *http.Request) {
	MoneyChart := app.NewChart()
	w.Header().Set("Content-Type", "application/json")
	data := map[string]any{
		"labels":   MoneyChart.GetLabels(),
		"datasets": MoneyChart.GetDatasets(),
	}

	json.NewEncoder(w).Encode(data)
}
