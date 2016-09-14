package rest

import (
	"encoding/json"
	"fmt"
	"github.com/anujkaushik-mfl/penguin/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

func StartRestServer() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Endpoints = [/images, /gif]")
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	fileNames, _ := filepath.Glob(config.IMAGE_STORAGE_TEMP + "*." + config.IMAGE_EXTENTION)
	images := fileNames
	if err := json.NewEncoder(w).Encode(images); err != nil {
		log.Fatal("Some error")
		panic(err)
	}
}

func GetGif(w http.ResponseWriter, r *http.Request) {
	fileNames, _ := filepath.Glob(config.IMAGE_STORAGE_TEMP + "*.gif")
	images := fileNames
	if err := json.NewEncoder(w).Encode(images); err != nil {
		log.Fatal("Some error")
		panic(err)
	}
}
