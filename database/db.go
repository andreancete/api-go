package database

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

var Db *sql.DB

//func que cria a conexão e retorna o BD
func ConectaFirebird() {
	var err error
	Db, err = sql.Open("firebirdsql", GetPathConexao())
	if err != nil {
		panic("Falha ao criar conexão! " + err.Error())
	}
}

func TestePing() bool {
	err := Db.Ping()
	if err != nil {
		panic("Falha ao conetar no banco! " + GetPathConexao() + err.Error())
	}
	return err == nil
}

//func que recebe o db e fecha a conexão
func DesconectaFirebird() {
	defer Db.Close()
}
