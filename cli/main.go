package main

import (
	"fmt"
	"log"
	"os"

	"github.com/timurgen/sesam-golang-client/v0/api"
)

func main() {
	node := os.Getenv("SESAM_NODE")
	jwt := os.Getenv("SESAM_JWT")

	if node == "" || jwt == "" {
		log.Println("[ERROR] Didn't find required arguments")
		os.Exit(2)
	}

	api := api.NewAPI(node, jwt)
	vars, err := api.GetEnvironmentVariables()
	if err != nil {
		log.Printf("Oppbs")
		os.Exit(2)
	}
	fmt.Printf("%v", vars)

}
