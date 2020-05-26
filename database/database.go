package database

import "database/sql"

//InitDB faz a conexão com o banco de dados MariaDB
func InitDB() *sql.DB {
	connectionString := "root:dani2016@tcp(localhost:3306)/marketplace"

	databaseConnection, err := sql.Open("mysql", connectionString)
	defer databaseConnection.Close()
	if err != nil {
		panic(err.Error)
	}
	return databaseConnection
}
