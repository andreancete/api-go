package database

import (
	"database/sql"
	"evcont/util"
)

// funcao que conecta no banco roda sql e fecha conexão
func Query(query string, params ...interface{}) (*sql.Rows, error) {
	ConectaFirebird()
	rows, err := Db.Query(query, params...)
	//defer rows.Close()
	util.TrataErro("SQL:"+query, err)
	DesconectaFirebird()
	util.TrataErro("Erro de conexão com o banco", err)
	return rows, err
}

func QueryRow(query string, params ...interface{}) (*sql.Row) {
	ConectaFirebird()
	row := Db.QueryRow(query, params...)
	DesconectaFirebird()
	return row
}
