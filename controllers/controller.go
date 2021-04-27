package controllers

import (
	"evcont/models"
	"evcont/repository"
	"evcont/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type DataEvcont struct {
	Router     *mux.Router
	Repository *repository.DatabaseRepository
	Server     http.Server
}

func NewCore() *DataEvcont {
	r := mux.NewRouter()
	repo := repository.NewRepository()
	s := http.Server{
		ErrorLog:     log.New(os.Stdout, "DataEvcont ", log.LstdFlags),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return &DataEvcont{Router: r, Repository: repo, Server: s}
}

func (D *DataEvcont) Ola(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("{\"api\": \"Rodando\"}"))
}

func (D *DataEvcont) FetchEventos(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("FechEventos")
	vars := mux.Vars(r)
	var eventos []models.Evento
	page, _ := strconv.Atoi(vars["pageNumber"])
	orderBy := vars["order"]
	uf := vars["uf"]
	ano := vars["ano"]
	fmt.Println("dados:", page, orderBy, uf, ano)

	if len(vars["uf"]) > 0 && len(vars["ano"]) > 0 {
		eventos = D.Repository.GetEventos(uf, ano, page, orderBy)
	}
	util.Respond(rw, map[string]interface{}{"eventos": eventos})
}

func (D *DataEvcont) FetchGrupos(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("FechGrupos")
	vars := mux.Vars(r)
	var grupos []models.Grupo

	page, _ := strconv.Atoi(vars["pageNumber"])
	eventoId, _ := vars["evento"]
	orderBy := vars["order"]
	uf := vars["uf"]
	ano := vars["ano"]

	if len(vars["uf"]) > 0 && len(vars["ano"]) > 0 {
		grupos = D.Repository.GetGrupos(uf, ano, eventoId, page, orderBy)
	}
	util.Respond(rw, map[string]interface{}{"grupos": grupos})
	//_ = json.NewEncoder(rw).Encode(grupos)
}

func GetPortaServidor() string {
	return repository.GetPortaServidor()
}

func (D *DataEvcont) GetVersao(rw http.ResponseWriter, r *http.Request) {
	versao := D.Repository.GetVersao()
	//_ = json.NewEncoder(rw).Encode(versao)
	util.Respond(rw, map[string]interface{}{"versao": versao})
}
