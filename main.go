package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/khdoba2000/REST-API-bookstore/urls"

)


func main() {

	// insert

	// hardcoded
	// for i := 3; i <= 10; i++ {
	// 	insertStmt := fmt.Sprintf(`insert into "students"("id", "name", "level") values(%v, %v, 2)`, i, i)
	// 	_, e := db.Exec(insertStmt)
	// 	CheckError(e)
	// }
	// // dynamic
	// insertDynStmt := `insert into "students"("id", "name", "level") values($1, $2, 3)`
	// for i := 300; i <= 335; i++ {
	// 	_, e := db.Exec(insertDynStmt, i, fmt.Sprintf("student%v", i))
	// 	CheckError(e)
	// }

	// updateStmt := `update "students" set "name"=$2 where "name"=$1`
	// _, e := db.Exec(updateStmt, "sName", "StudentName")
	// CheckError(e)

	// // Delete
	// deleteStmt := `delete from "students" where id=$1`
	// _, e := db.Exec(deleteStmt, 6)
	// CheckError(e)

	// rows, err := db.Query(`SELECT "id", "name", "level" FROM "students"`)
	// CheckError(err)

	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var level int

	// 	err = rows.Scan(&id, &name, &level)
	// 	CheckError(err)
	// 	fmt.Println(id, name, level)
	// }

	// CheckError(err)
	urls.handleRequest()
}
