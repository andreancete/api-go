package repository

import (
	"evcont/database"
	"evcont/models"
)

//var grupos []models.Grupo

func (r *DatabaseRepository) GetGrupos(uf, ano string, eventoId string, page int, order string) []models.Grupo {
	//n, n1 := util.PageLength(page)
	// if grupos == nil {
	var grupos []models.Grupo
	rows, err := database.Query(models.SqlGrupo, uf, ano)
	grupos, err = models.ScanGrupo(rows)
	if err != nil {
		grupos = nil
	}
	return grupos
}
