package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"api-go/src/routes"
)

func handleRequests() {
	http.HandleFunc("/", routes.Home)

	err := godotenv.Load(".env")

	port := os.Getenv("PORT")

  if err != nil {
    port = "3800"
	}
	
	log.Println("server listend on port " + port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func main() {
	handleRequests()
}