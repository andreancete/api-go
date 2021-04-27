package router

import (
	"evcont/controllers"
)

func RegisterRoutes(core *controllers.DataEvcont) {
	core.Router.HandleFunc("/api", core.Ola).Methods("GET")
	core.Router.HandleFunc("/api/eventos", core.FetchEventos).Methods("GET").Queries("uf", "{uf}", "ano", "{ano}")
	core.Router.HandleFunc("/api/grupos", core.FetchGrupos).Methods("GET").Queries("uf", "{uf}", "ano", "{ano}")
	core.Router.HandleFunc("/api/versao", core.GetVersao).Methods("GET")
}
