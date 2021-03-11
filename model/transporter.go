package model

type Transporter struct {
	Oferta           int     `json:"oferta"`
	Cnpj             string  `json:"cnpj"`
	Logotipo         string  `json:"logotipo"`
	Nome             string  `json:"nome"`
	Servico          string  `json:"servico"`
	PrazoEntrega     int     `json:"prazo_entrega"`
	EntregaEstimada  string  `json:"entrega_estimada"`
	DescricaoServico string  `json:"descricao_servico"`
	Validade         string  `json:"validade"`
	CustoFrete       float64 `json:"custo_frete"`
	PrecoFrete       float64 `json:"preco_frete"`
}

type TransporterOffer struct {
	TokenOferta     string        `json:"token_oferta"`
	Transportadoras []Transporter `json:"transportadoras"`
}
