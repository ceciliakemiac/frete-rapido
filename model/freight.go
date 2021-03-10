package model

import "gorm.io/gorm"

type Freight struct {
	gorm.Model
	TransporterCNPJ string  `json:"transporter_cnpj" gorm:"not null;default:null"`
	Transporter     string  `json:"transporter_name" gorm:"not null;default:null"`
	Price           float64 `json:"price" gorm:"not null;default:null"`
	QuoteID         uint    `json:"quote_id" gorm:"not null;default:null"`
}
