package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"

	m "github.com/noliva/api-maps/models"
)

func CreateMap(w http.ResponseWriter, req *http.Request) {
	var newAreaMap m.MyMap
	_ = json.NewDecoder(req.Body).Decode(&newAreaMap)
	newAreaMap.ID = uuid.NewV4().String()
	newAreaMap.Insert()
	json.NewEncoder(w).Encode(newAreaMap)
}

func GetMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	newAreaMap := m.FindMapById(params["id"])

	json.NewEncoder(w).Encode(&newAreaMap)
}

func GetMaps(w http.ResponseWriter, req *http.Request) {
	areaMaps := m.FindAllMaps()

	json.NewEncoder(w).Encode(&areaMaps)
}