package models

type ConfiguracaoDB struct {
	Banco banco
	Servidor servidor
}

type banco struct {
	Host    string `json:"Host,omitempty"`
	Porta    string `json:"Port,omitempty"`
	Arquivo string `json:"Path,omitempty"`

}

type servidor struct {
	Porta    string `json:"Porta,omitempty"`
}


