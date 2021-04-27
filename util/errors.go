package util

import (
	"errors"
	"fmt"
)

type erros struct {
	titulo    string
	descricao string
}

var erro erros

func TrataErro(msg string, err error) {
	if err != nil {
		erro.titulo = msg
		erro.descricao = err.Error()
		fmt.Println(msg + " :Erro " + err.Error())
		return
	} else {
		erro.titulo = ""
		erro.descricao = ""
	}

}

var (
	NotFound            = errors.New("Request not found")
	InternalServerError = errors.New("Internal Server Error")
)
