package models

import (
	"database/sql"
	"evcont/util"
)

const SqlEvento = `SELECT UF,ANO,EVENTO,NOME,LANCAMENTO,TIPO_LANC,CODIGOAUDESP,CAST(OBS AS VARCHAR(500)) OBS,CAST(FILTRO AS VARCHAR(500)) FILTRO,EVENTO_TCE,SEQ_EVENTO,MES 
										FROM EVENTO 
										WHERE UF=? and ano=? ORDER BY EVENTO`

type Evento struct {
	UF             string  `json:"uf"`
	Ano            string  `json:"ano"`
	Evento         string  `json:"evento"`
	Nome           string  `json:"nome"`
	Lancamento     string  `json:"lancamento"`
	TipoLancamento string  `json:"tipoLancamento"`
	CodigoAudesp   string  `json:"codigoAudesp"`
	Obs            string  `json:"obs"`
	Filtro         string  `json:"filtro"`
	EventoTce      string  `json:"eventoTce"`
	Seq_Evento     string  `json:"seqEvento"`
	Mes            string  `json:"mes"`
	Grupos         []Grupo `json:"grupos,omitempty"`
}

func ScanEvento(rows *sql.Rows) []Evento {
	var ev []Evento
	evento := Evento{}
	var evuf, evano, evt, evtNome, evtLancamento, evCodigoAudesp, evtTipo, evtTce, evtSeq, evmes sql.NullString
	var evfiltro, evobs sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&evuf,
			&evano,
			&evt,
			&evtNome,
			&evtLancamento,
			&evtTipo,
			&evCodigoAudesp,
			&evobs,
			&evfiltro,
			&evtTce,
			&evtSeq,
			&evmes,
		)
		evento.UF = evuf.String
		evento.Ano = evano.String
		evento.Evento = evt.String
		evento.Nome = evtNome.String
		evento.Lancamento = evtLancamento.String
		evento.TipoLancamento = evtTipo.String
		evento.CodigoAudesp = evCodigoAudesp.String
		evento.Obs = evobs.String
		evento.Filtro = evfiltro.String
		evento.EventoTce = evtTce.String
		evento.Seq_Evento = evtSeq.String
		evento.Mes = evmes.String

		util.TrataErro("ScanEvento:", err)
		ev = append(ev, evento)
	}
	return ev
}
