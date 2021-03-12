package database

import (
	"time"

	"github.com/ceciliakemiac/frete-rapido/model"
)

func (db *Database) CreateFreight(transporterOffers *model.TransporterOffer) (*model.Freights, error) {
	now := time.Now()
	quote := model.Quote{
		CreatedAt: now,
	}

	createdQuote, err := db.createQuote(&quote)
	if err != nil {
		return nil, err
	}

	var freights []model.Freight
	transporters := transporterOffers.Transportadoras
	for _, transporter := range transporters {
		freight := model.Freight{
			Nome:         transporter.Nome,
			Servico:      transporter.Servico,
			PrazoEntrega: transporter.PrazoEntrega,
			PrecoFrete:   transporter.PrecoFrete,
			QuoteID:      createdQuote.ID,
		}

		freights = append(freights, freight)
	}

	if err := db.PG.Create(&freights).Error; err != nil {
		return nil, err
	}

	transportadoras := model.Freights{
		Transportadoras: freights,
	}

	return &transportadoras, nil
}

func (db *Database) createQuote(quote *model.Quote) (*model.Quote, error) {
	if err := db.PG.Create(&quote).Error; err != nil {
		return nil, err
	}

	return quote, nil
}
