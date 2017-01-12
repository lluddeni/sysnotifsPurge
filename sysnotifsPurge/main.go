package main

import (
	"fmt"
	"log"
	"net/http"
	"sysnotifsPurge/routes"
)

func main() {

	fmt.Println("main entry")
	//Setup the API routes
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
