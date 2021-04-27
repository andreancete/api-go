package repository

import (
	"evcont/database"
	"evcont/models"
	"evcont/util"
)

func (r *DatabaseRepository) GetVersao() models.Versao {
	var versao models.Versao
	var valor int
	err := database.QueryRow(models.SqlVersao).Scan(&valor)
	util.TrataErro("Executando a query GETVERSAO", err)
	versao.Numero = valor
	//versao.Numero, err = strconv.Atoi(valorStr)
	//versao.Numero =  models.ScanVersao(rows)
	return versao
}

// func ExampleDB_QueryContext() {
// 	var ctx context.Context
// 	rows, err := database.Db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	names := make([]string, 0)

// 	for rows.Next() {
// 		var name string
// 		if err := rows.Scan(&name); err != nil {
// 			// Check for a scan error.
// 			// Query rows will be closed with defer.
// 			log.Fatal(err)
// 		}
// 		names = append(names, name)
// 	}
// 	// If the database is being written to ensure to check for Close
// 	// errors that may be returned from the driver. The query may
// 	// encounter an auto-commit error and be forced to rollback changes.
// 	rerr := rows.Close()
// 	if rerr != nil {
// 		log.Fatal(rerr)
// 	}

// 	// Rows.Err will report the last error encountered by Rows.Scan.
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s are %d years old", strings.Join(names, ", "), age)
// }
