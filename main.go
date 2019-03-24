package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
	fmt.Println()}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/{id}", getHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
