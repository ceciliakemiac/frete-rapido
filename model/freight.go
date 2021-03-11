package model

type Freight struct {
	ID           int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome         string  `json:"nome" gorm:"not null;default:null"`
	Servico      string  `json:"servico" gorm:"not null;default:null"`
	PrazoEntrega int     `json:"prazo_entrega" gorm:"not null;default:null"`
	PrecoFrete   float64 `json:"preco_frete" gorm:"not null;default:null"`
	QuoteID      int     `json:"quote_id" gorm:"not null;default:null"`
}
