package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// update
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")

	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	// delete
	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt2.Exec(3)
}
