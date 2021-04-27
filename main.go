package main

import (
	"evcont/controllers"
	"evcont/router"
	"evcont/util"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	porta := controllers.GetPortaServidor()
	core := controllers.NewCore()
	router.RegisterRoutes(core)
	handler := cors.Default().Handler(core.Router)
	err := http.ListenAndServe(":"+porta, handler)
	util.TrataErro("Erro iniciando o servidor:", err)

}
