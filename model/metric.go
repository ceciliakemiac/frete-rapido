package model

type Metric struct {
	Transportadora   string  `json:"transportadora"`
	TotalOcorrencias int     `json:"total_ocorrencias"`
	TotalPrecos      float64 `json:"total_precos"`
	MediaPrecos      float64 `json:"media_precos"`
}

type ValueFreight struct {
	Nome         string  `json:"nome"`
	Servico      string  `json:"servico"`
	PrazoEntrega int     `json:"prazo_entrega"`
	Valor        float64 `json:"valor"`
}

type Metrics struct {
	Fretes          []Metric     `json:"fretes"`
	FreteMaisCaro   ValueFreight `json:"frete_mais_caro"`
	FreteMaisBarato ValueFreight `json:"frete_mais_barato"`
}
