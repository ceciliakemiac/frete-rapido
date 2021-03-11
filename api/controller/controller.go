package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ceciliakemiac/frete-rapido/api/service"
	"github.com/ceciliakemiac/frete-rapido/model"
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

	basePath.Path("/quote").HandlerFunc(c.PostQuote).Methods(http.MethodPost)

	return c
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello from Frete Rápido Desafio backend server!")
	w.Write(res)
}

func (c *Controller) PostQuote(w http.ResponseWriter, r *http.Request) {
	var volumeData model.VolumeData

	if err := json.NewDecoder(r.Body).Decode(&volumeData); err != nil {
		log.Println("Error decoding volumes")
		http.Error(w, "Error decoding volumes", http.StatusInternalServerError)
		return
	}

	remetente := &model.Remetente{
		Cnpj: os.Getenv("CNPJ"),
	}

	volumeSecureData := &model.VolumeSecureData{
		Remetente:        *remetente,
		Destinatario:     volumeData.Destinatario,
		Volumes:          volumeData.Volumes,
		Token:            os.Getenv("TOKEN"),
		CodigoPlataforma: os.Getenv("CODIGO_PLATAFORMA"),
	}

	volumesJson, err := json.Marshal(volumeSecureData)
	if err != nil {
		log.Println("Error json.Marshal")
	}

	transporterOffers, err := getOffersData(volumesJson)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting offers", http.StatusInternalServerError)
		return
	}

	freights, err := c.service.CreateFreight(transporterOffers)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(freights)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func getOffersData(volumes []byte) (*model.TransporterOffer, error) {
	client := &http.Client{}
	url := os.Getenv("QUOTE_SIMULATOR_URL")

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(volumes))
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var transporterOffer model.TransporterOffer
	if err = json.NewDecoder(res.Body).Decode(&transporterOffer); err != nil {
		return nil, err
	}

	return &transporterOffer, nil
}
