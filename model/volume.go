package model

type Volume struct {
	Tipo           int     `json:"tipo" gorm:"not null;default:null"`
	Quantidade     int     `json:"quantidade" gorm:"not null;default:null"`
	Peso           int     `json:"peso" gorm:"not null;default:null"`
	Valor          int     `json:"valor" gorm:"not null;default:null"`
	Sku            string  `json:"sku"`
	Tag            string  `json:"tag"`
	Descricao      string  `json:"descricao"`
	Altura         float64 `json:"altura" gorm:"not null;default:null"`
	Largura        float64 `json:"largura" gorm:"not null;default:null"`
	Comprimento    float64 `json:"comprimento" gorm:"not null;default:null"`
	VolumesProduto int     `json:"volumes_produto"`
	Consolidar     bool    `json:"consolidar"`
	Sobreposto     bool    `json:"sobreposto"`
	Tombar         bool    `json:"tombar"`
}

type Endereco struct {
	Cep string `json:"cep" gorm:"not null;default:null"`
}

type Destinatario struct {
	Endereco Endereco `json:"endereco" gorm:"not null;default:null"`
}

type Remetente struct {
	Cnpj string `json:"cnpj" gorm:"not null;default:null"`
}

type Expedidor struct {
	Cnpj     string   `json:"cnpj"`
	Endereco Endereco `json:"endereco"`
}

type VolumeData struct {
	Destinatario Destinatario `json:"destinatario" gorm:"not null;default:null"`
	Volumes      []Volume     `json:"volumes" gorm:"not null;default:null"`
}

type VolumeSecureData struct {
	Remetente        Remetente    `json:"remetente" gorm:"not null;default:null"`
	Destinatario     Destinatario `json:"destinatario" gorm:"not null;default:null"`
	Volumes          []Volume     `json:"volumes" gorm:"not null;default:null"`
	Token            string       `json:"token" gorm:"not null;default:null"`
	CodigoPlataforma string       `json:"codigo_plataforma" gorm:"not null;default:null"`
}
