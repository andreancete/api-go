package models

import (
	"database/sql"
	"evcont/util"
)

const SqlGrupo = `SELECT UF,ANO,EVENTO,GRUPO,CONTA,NOME,D_C,SQL,DESDOBRAR,DATA,LANCA,ISF_GRUPO 
								FROM GRUPO 
								WHERE UF=? and ANO=? ORDER BY EVENTO,GRUPO`

type Blob []byte

type Grupo struct {
	UF        string `json:"uf"`
	Ano       string `json:"ano"`
	Evento    string `json:"evento"`
	Grupo     string `json:"erupo"`
	Conta     string `json:"conta"`
	Nome      string `json:"nome"`
	DC        string `json:"dc"`
	Sql       Blob   `json:"sql"`
	Desdobrar string `json:"desdobrar"`
	Data      string `json:"data"`
	Lanca     string `json:"lanca"`
	Isf       string `json:"isf"`
}

func ScanGrupo(rows *sql.Rows) ([]Grupo, error) {
	var gr []Grupo
	var err error
	grupo := Grupo{}
	var gruf, grano, grEvento, grGrupo, grConta, grNome, grDC, grDesdobrar, grData, grLanca, grIsf sql.NullString
	var grSql Blob
	for rows.Next() {
		err := rows.Scan(
			&gruf,
			&grano,
			&grEvento,
			&grGrupo,
			&grConta,
			&grNome,
			&grDC,
			&grSql,
			&grDesdobrar,
			&grData,
			&grLanca,
			&grIsf,
		)
		grupo.UF = gruf.String
		grupo.Ano = grano.String
		grupo.Evento = grEvento.String
		grupo.Grupo = grGrupo.String
		grupo.Conta = grConta.String
		grupo.Nome = grNome.String
		grupo.DC = grDC.String
		grupo.Sql = grSql
		grupo.Desdobrar = grDesdobrar.String
		grupo.Data = grData.String
		grupo.Lanca = grLanca.String
		grupo.Isf = grIsf.String

		util.TrataErro("ScanGrupo:", err)

		gr = append(gr, grupo)
	}
	return gr, err
}
