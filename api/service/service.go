package service

import (
	"time"

	"github.com/ceciliakemiac/frete-rapido/model"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateFreight(transporterOffers *model.TransporterOffer) (*[]model.Freight, error) {
	now := time.Now()
	quote := model.Quote{
		CreatedAt: now,
	}

	createdQuote, err := s.createQuote(&quote)
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

	if err := s.db.Create(&freights).Error; err != nil {
		return nil, err
	}

	return &freights, nil
}

func (s *Service) createQuote(quote *model.Quote) (*model.Quote, error) {
	if err := s.db.Create(&quote).Error; err != nil {
		return nil, err
	}

	return quote, nil
}
