package main

import (
	"GolangRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databaseConnection := database.InitDB()

	fmt.Println(databaseConnection)
	/*	r := chi.NewRouter()
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome"))

		})
		http.ListenAndServe(":3000", r)
	*/

}
