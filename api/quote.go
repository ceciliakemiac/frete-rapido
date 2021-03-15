package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	remetente := model.Remetente{
		Cnpj: os.Getenv("CNPJ"),
	}

	volumeSecureData, err := getVolumeSecureData(&volumeData, &remetente)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
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

func getVolumeSecureData(volume *model.VolumeData, remetente *model.Remetente) (*model.VolumeSecureData, error) {
	if volume == nil || remetente == nil {
		return nil, fmt.Errorf("GetVolumeSecureData: Volume or Remetente are invalid")
	}

	volumeSecureData := &model.VolumeSecureData{
		Remetente:        *remetente,
		Destinatario:     volume.Destinatario,
		Volumes:          volume.Volumes,
		Token:            os.Getenv("TOKEN"),
		CodigoPlataforma: os.Getenv("CODIGO_PLATAFORMA"),
	}

	return volumeSecureData, nil
}

func getOffersData(volumes []byte) (*model.TransporterOffer, error) {
	if volumes == nil || len(volumes) == 0 {
		return nil, fmt.Errorf("GetOffersData: Volumes are nil or invalid")
	}

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
