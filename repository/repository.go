package repository

import (
	"database/sql"
	"evcont/database"
	"evcont/models"
)

type Repository interface {
	GetEventos(uf, ano string) []models.Evento
	GetGrupos(uf, ano string) []models.Grupo
	GetVersao() []models.Versao
}

type DatabaseRepository struct {
	DB *sql.DB
}

func NewRepository() *DatabaseRepository {
	database.ConectaFirebird()
	database.DesconectaFirebird()
	return &DatabaseRepository{DB: database.Db}

}

func GetPortaServidor() string {
	return database.CarregaParametrosServer()
}
