package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"

	"github.com/ceciliakemiac/frete-rapido/api/controller"
)

type Server struct {
	addr       string
	router     *mux.Router
	db         *gorm.DB
	controller *controller.Controller
}

func NewServer(addr string, db *gorm.DB) (*Server, error) {
	s := &Server{
		addr:   addr,
		db:     db,
		router: mux.NewRouter(),
	}

	s.controller = controller.NewController(s.router, s.db)

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
