package main

import (
	"fmt"
	"log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "github.com/garrettreed/garrettreed.info/api/aggregate"
)

func getSiteData(w http.ResponseWriter, r *http.Request) {
    siteData, siteDataErr := aggregate.GetAggregateData()
    if siteDataErr != nil {
        log.Fatal(siteDataErr)
    }
    json.NewEncoder(w).Encode(siteData)
}

func main() {
    envErr := godotenv.Load("../.env")
    if envErr != nil {
        log.Fatal(envErr)
    }

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/data", getSiteData)
	port := 8080
	fmt.Print("Listening on port ", port, "\n")

    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
