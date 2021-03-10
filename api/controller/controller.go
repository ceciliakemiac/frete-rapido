package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ceciliakemiac/frete-rapido/api/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Controller struct {
	router  *mux.Router
	db      *gorm.DB
	service *service.Service
}

func NewController(router *mux.Router, db *gorm.DB) *Controller {
	c := &Controller{
		router:  router,
		db:      db,
		service: service.NewService(db),
	}

	basePath := c.router.PathPrefix("/api").Subrouter()
	basePath.Path("").HandlerFunc(c.Ping).Methods(http.MethodGet)

	return c
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello from Frete Rápido Desafio backend server!")
	w.Write(res)
}
