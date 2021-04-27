package models

// import (
// 	"database/sql"
// 	"evcont/util"
// )

type Receita struct {
	UF                 string `json:"uf"`
	ANO                string `json:"ano"`
	CODRE              string `json:"codre"`
	DESCRICAO          string `json:"descricao"`
	TIPO               string `json:"tipo"`
	N1                 string `json:"n1"`
	N2                 string `json:"n2"`
	N3                 string `json:"n3"`
	N4                 string `json:"n4"`
	N5                 string `json:"n5"`
	N6                 string `json:"n6"`
	N7                 string `json:"n7"`
	N8                 string `json:"n8"`
	CONTA_VARIACAO     string `json:"contaVariacao"`
	FILTRO_VARIACAO    Blob   `json:"filtroVariacao"`
	FILTRO_COMPETEN    Blob   `json:"filtroCompetencia"`
	CONTA_COMPETEN     string `json:"contaCompeten"`
	TIPO_FICHA         string `json:"tipoFicha"`
	UTILIZA            string `json:"utiliza"`
	DESCR_ABREVIADO    string `json:"descrAbreviado"`
	RETIRA_AUDESP      string `json:"retiraAudesp"`
	PERMITE_DESDOBRO   string `json:"permiteDesdobro"`
	FUNCAO             string `json:"funcao"`
	DESTINACAO_LEGAL   string `json:"destinacaoLegal"`
	SUPORTE_DOCUMENTAL string `json:"suportedocumental"`
	AMPARO_LEGAL       string `json:"amparoLegal"`
	NIVEL              int    `json:"nivel"`
	PERMITE_DEDUCAO    string `json:"permiteDeducao"`
	EXCLUIDO           string `json:"excluido"`
}

// func ScanReceita(rows *sql.Rows) ([]Receita, error) {
// 	var rec []Receita
// 	var err error
// 	receita := Receita{}
// 	var recuf, grano, grEvento, grGrupo, grConta, grNome, grDC, grDesdobrar, grData, grLanca, grIsf sql.NullString
// 	var grSql Blob
// 	for rows.Next() {
// 		err := rows.Scan(
// 			&gruf,
// 			&grano,
// 			&grEvento,
// 			&grGrupo,
// 			&grConta,
// 			&grNome,
// 			&grDC,
// 			&grSql,
// 			&grDesdobrar,
// 			&grData,
// 			&grLanca,
// 			&grIsf,
// 		)
// 		grupo.UF = gruf.String
// 		grupo.Ano = grano.String
// 		grupo.Evento = grEvento.String
// 		grupo.Grupo = grGrupo.String
// 		grupo.Conta = grConta.String
// 		grupo.Nome = grNome.String
// 		grupo.DC = grDC.String
// 		grupo.Sql = grSql
// 		grupo.Desdobrar = grDesdobrar.String
// 		grupo.Data = grData.String
// 		grupo.Lanca = grLanca.String
// 		grupo.Isf = grIsf.String

// 		util.TrataErro("ScanGrupo:", err)

// 		gr = append(gr, grupo)
// 	}
// 	return gr, err
// }
