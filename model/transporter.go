package model

type Transporter struct {
	Oferta           int     `json:"oferta" gorm:"not null;default:null"`
	Cnpj             string  `json:"cnpj" gorm:"not null;default:null"`
	Logotipo         string  `json:"logotipo" gorm:"not null;default:null"`
	Nome             string  `json:"nome" gorm:"not null;default:null"`
	Servico          string  `json:"servico" gorm:"not null;default:null"`
	PrazoEntrega     int     `json:"prazo_entrega" gorm:"not null;default:null"`
	EntregaEstimada  string  `json:"entrega_estimada" gorm:"not null;default:null"`
	DescricaoServico string  `json:"descricao_servico"`
	Validade         string  `json:"validade" gorm:"not null;default:null"`
	CustoFrete       float64 `json:"custo_frete" gorm:"not null;default:null"`
	PrecoFrete       float64 `json:"preco_frete" gorm:"not null;default:null"`
}

type TransporterOffer struct {
	TokenOferta     string        `json:"token_oferta" gorm:"not null;default:null"`
	Transportadoras []Transporter `json:"transportadoras" gorm:"not null;default:null"`
}
