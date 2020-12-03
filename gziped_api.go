package main

import (
	"encoding/json"
	"compress/gzip"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type AppSettings struct {
	COMPRESS bool
}

var (
	response Response
	settings AppSettings
)

func main() {
	settings = AppSettings{
		COMPRESS: true,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", compressedHome)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func compressedHome(res http.ResponseWriter, req *http.Request) {
	response = Response{
		Message: "This message is compressed",
	}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type","application/json")
	if settings.COMPRESS {
		res.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(res)

		_ = json.NewEncoder(gz).Encode(response)

		_ = gz.Close()
	} else {
		response.Message = "This message is not compressed"
		_ = json.NewEncoder(res).Encode(response)
	}
	res.Header().Set("Content-Type","application/json")




}
