package main

import (
	"fmt"
	"os"
)

func GetEnvOrPanic(name string) string {

	value, defined := os.LookupEnv(name)
	if !defined {
		panic(fmt.Sprintf("%s env var is not set", name))
	}

	return value

}
