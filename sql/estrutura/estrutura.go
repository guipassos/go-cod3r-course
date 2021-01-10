package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // colocado um _ porque não vamos usar ele explicitamente
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:1234@/") // "@/" se conectará com a raiz (sem especificar um database) pois nesse exemplo será criado o schema e a tabela
	if err != nil {
		panic(err)
	}
	defer db.Close() // defer: será fechado a conexão quando terminar essa função main
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY(id)
	)`)
}
