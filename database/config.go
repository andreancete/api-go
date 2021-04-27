package database

import (
	"encoding/json"
	"evcont/models"
	"evcont/util"
	"fmt"
	"io/ioutil"
	"os"
)

var Strconexao string
var Objconfiguracao models.ConfiguracaoDB

const nomearquivo = `caminhodb.json`

func carregaConfig(jsonFile *os.File) models.ConfiguracaoDB {
	//Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	//Declaração abreviada de um objeto do tipo
	objConfigDB := models.ConfiguracaoDB{}

	//Conversão da variável byte em um objeto do tipo struct
	json.Unmarshal(byteValueJSON, &objConfigDB)
	return objConfigDB
}

func salvaConfig(objconfig models.ConfiguracaoDB) {
	byteValueJSON, err := json.Marshal(objconfig)
	if err != nil {
		fmt.Println(err)
	}

	//Por fim vamos salvar em um arquivo JSON alterado.
	err = ioutil.WriteFile(nomearquivo, byteValueJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func carregapadrao() models.ConfiguracaoDB {
	objconf := models.ConfiguracaoDB{}
	objconf.Banco.Host = "localhost"
	objconf.Banco.Porta = "3055"
	objconf.Banco.Arquivo = "C:/scpi_8/TABDEF/TCE_PCASP.FDB"
	objconf.Servidor.Porta = "9090"
	return objconf
}

func CarregaParametrosServer() string {
	jsonFile, err := os.Open(nomearquivo)
	if err != nil {
		jsonFile, err = os.Create(nomearquivo)
		util.TrataErro("Erro ao criar arquivo "+nomearquivo, err)
	}
	conf := carregaConfig(jsonFile)
	if conf.Servidor.Porta == "" {
		conf = carregapadrao()
		salvaConfig(conf)
	}
	defer jsonFile.Close()
	fmt.Println()
	fmt.Printf("Servidor rodando na porta:%s", conf.Servidor.Porta)
	fmt.Println()
	fmt.Printf("Acesse localhost:%s/api", conf.Servidor.Porta)
	return conf.Servidor.Porta
}

func GetPathConexao() string {
	if Strconexao == "" {
		carregaParametrosDB()
		//"FSCSCPI8:scpi@localhost:3055/C:/scpi_8/TABDEF/TCE_PCASP.FDB"
		Strconexao = "FSCSCPI8:scpi@" + Objconfiguracao.Banco.Host + ":" + Objconfiguracao.Banco.Porta + "/" + Objconfiguracao.Banco.Arquivo
		fmt.Println()
		fmt.Printf("conectado em :%s", Strconexao)
	}
	return Strconexao
}

//carrega do arquivo as configurações do banco e atribui na variavel Objconfiguracao
func carregaParametrosDB() {
	jsonFile, err := os.Open(nomearquivo)
	if err != nil {
		jsonFile, err = os.Create(nomearquivo)
		util.TrataErro("Erro ao criar arquivo "+nomearquivo, err)
	}
	conf := carregaConfig(jsonFile)
	if conf.Banco.Host == "" {
		conf = carregapadrao()
		salvaConfig(conf)
	}
	defer jsonFile.Close()
	Objconfiguracao = conf

}
