package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type maps struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Boundaries  string `json:"boundaries"`
	GroupOwner  string `json:"groupOwner"`
}

var mapa []maps

func createMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var mapa1 maps
	_ = json.NewDecoder(req.Body).Decode(&mapa1)
	mapa1.ID = params["id"]
	mapa = append(mapa, mapa1)
	json.NewEncoder(w).Encode(mapa)
}

func getMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range mapa {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&maps{})
}

func getMaps(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mapa)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mapa)
}

func middlewareFunction(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	mapa = append(mapa, maps{ID: "uuid", Name: "WC1-1", Description: "just a description", Boundaries: "boundaries", GroupOwner: "owner"})
	mapa = append(mapa, maps{ID: "4232", Name: "NW1-3", Description: "somewhere over the rainbow", Boundaries: "boundaries x to yq"})

	router := mux.NewRouter()

	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/maps", getMaps).Methods("GET")
	router.HandleFunc("/maps/{id}", getMap).Methods("GET")
	router.HandleFunc("/maps/{id}", createMap).Methods("POST")
	router.HandleFunc("/maps/{id}", createMap).Methods("DELETE")

	http.Handle("/", middlewareFunction(router))

	log.Fatal(http.ListenAndServe(":12345", router))
}
