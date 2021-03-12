package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ceciliakemiac/frete-rapido/database"
)

type Server struct {
	addr   string
	router *mux.Router
	db     *database.Database
}

func NewServer(addr string, db *database.Database) (*Server, error) {
	s := &Server{
		addr:   addr,
		db:     db,
		router: mux.NewRouter(),
	}

	s.CreateRoutes(s.router, s.db)

	return s, nil
}

func (s *Server) Run() error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD"},
		AllowedHeaders: []string{"*"},
	})

	log.Println("Http Server starting to listen at", s.addr)
	err := http.ListenAndServe(s.addr, c.Handler(s.router))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) CreateRoutes(router *mux.Router, db *database.Database) {
	basePath := router.PathPrefix("/api").Subrouter()
	basePath.Path("").HandlerFunc(s.Ping).Methods(http.MethodGet)

	basePath.Path("/quote").HandlerFunc(s.CreateQuote).Methods(http.MethodPost)
	basePath.Path("/metrics").HandlerFunc(s.GetMetrics).Methods(http.MethodGet)
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello from Frete Rápido Desafio backend server!")
	w.Write(res)
}
