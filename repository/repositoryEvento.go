package repository

import (
	"evcont/database"
	"evcont/models"
	"evcont/util"
)



func (r *DatabaseRepository) GetEventos(uf, ano string, page int, order string) []models.Evento {
	// n, n1 := util.PageLength(page)
	//rows, err := database.Query(database.SelectEventos, uf, ano)
	// if eventos == nil {
	rows, err := database.Query(models.SqlEvento, uf, ano)
	util.TrataErro("Executando a query geteventos", err)
	eventos := models.ScanEvento(rows)

	// }
	return eventos

}
