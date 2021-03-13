package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/ceciliakemiac/frete-rapido/model"
)

func (s *Server) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var volumeData model.VolumeData

	if err := json.NewDecoder(r.Body).Decode(&volumeData); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
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
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transporterOffers, err := getOffersData(volumesJson)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	freights, err := s.db.CreateFreight(transporterOffers)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(freights)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getOffersData(volumes []byte) (*model.TransporterOffer, error) {
	client := &http.Client{Timeout: 10 * time.Second}
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
