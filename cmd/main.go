package main

import (
	"database/sql"
	"log"

	"github.com/AnirudhV16/UShort/cmd/api"
	"github.com/AnirudhV16/UShort/config"
	"github.com/AnirudhV16/UShort/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal()
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal()
	}

}

func initStorage(db *sql.DB) {
	//this actually connects database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("db is successfully connected!!")
}
