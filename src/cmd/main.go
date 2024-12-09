package main

import (
	"fmt"
	"log"

	"github.com/garrickedd/ReLibca/src/infrastructure"
)

func main() {
	database, err := infrastructure.Postgres()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(database)
}
