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

	if _, isDefined := os.LookupEnv("PG_ENABLED"); isDefined {
		TestPGConnection()
	}

	if _, isDefined := os.LookupEnv("WS_ENABLED"); isDefined {
		OpenWebservice()
	}

}
