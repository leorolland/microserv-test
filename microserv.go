package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Starting microserv instance...")

	start()

	fmt.Println("Ready !")

}

func start() {

	_, pgEnabled := os.LookupEnv("PG_ENABLED")
	var db *MicroservDB

	if pgEnabled {
		db = NewMicroservDB()
		if tableName, defined := os.LookupEnv("PG_TABLE"); defined {
			db.CreateTable(tableName)
		}
	}

	if _, isDefined := os.LookupEnv("WS_ENABLED"); isDefined {
		OpenWebservice(db)
	}

}
