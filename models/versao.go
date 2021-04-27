package models

import (
	"database/sql"
	"evcont/util"
	"strconv"
)

type Versao struct {
	Numero int `json:"numero"`
}

const SqlVersao = `SELECT NROVERSAO FROM VERSAO`

func ScanVersao(rows *sql.Rows) Versao {

	v := Versao{}
	var vers sql.NullString
	for rows.Next() {
		err := rows.Scan(
			&vers,
		)
		v.Numero, err = strconv.Atoi(vers.String)
		util.TrataErro("ScanVersao:", err)

	}
	return v
}
