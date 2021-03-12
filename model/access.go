package model

type Access struct {
	Cpnj             string `json:"cnpj"`
	Token            string `json:"token"`
	CodigoPlataforma string `json:"codigo_plataforma"`
}

type Block struct {
	Bloqueio bool `json:"bloqueio"`
	Saldo    int  `json:"saldo"`
}
