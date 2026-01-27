package main

import (
	"devbook-api/app/config"
	"devbook-api/app/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadConfig()

	r := router.Router()

	fmt.Printf("Rodando na porta %d", config.Port)
	fmt.Println(" ")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
