package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"D:/PERSONAL/GO/serrr/bookservice/bookservice"
)

func main() {
	router := mux.NewRouter()
	api.InitializeRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
