package models

import (
	"log"
	"github.com/noliva/api-maps/database"
)

type MyMap struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Boundaries  string    `json:"boundaries"`
	GroupOwner  string    `json:"groupOwner"db:"group_owner"`
	Addresses   []address `json:"addresses"`
}

func (m MyMap) Insert() {
	db := database.Connector()
	query := `INSERT INTO area_map (id, name, description, boundaries, group_owner) VALUES (?, ?, ?, ?, ?)`
	db.MustExec(query, m.ID, m.Name, m.Description, m.Boundaries, m.GroupOwner)
}

func FindMapById(id string) MyMap {
	db := database.Connector()
	var newAreaMap MyMap
	err := db.Select(&newAreaMap, "SELECT * FROM area_map WHERE id=?", id)

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return newAreaMap
}

func FindAllMaps() []MyMap {
	db := database.Connector()
	areaMaps := []MyMap{}
	err := db.Select(&areaMaps, "SELECT * FROM area_map")

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return areaMaps
}