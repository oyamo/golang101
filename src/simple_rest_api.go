package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

var (
	response Response
)



func main() {


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Homepage)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Homepage(w http.ResponseWriter, request *http.Request) {
	response = Response{
		Message: "Hello " + request.RemoteAddr + " we got a " + request.Method + " from you",
	}
	w.Header().Add("content-type", "application/json")
	w.Header().Add("x-server", "Golang XServer")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}