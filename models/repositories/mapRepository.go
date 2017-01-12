package repositories

import (
	"github.com/noliva/api-maps/database"
	"log"
	"github.com/noliva/api-maps/models"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db = database.Connector()
}

func Insert(m models.MyMap) {
	query := `INSERT INTO area_map (id, name, description, boundaries, group_owner) VALUES (?, ?, ?, ?, ?)`
	db.MustExec(query, m.ID, m.Name, m.Description, m.Boundaries, m.GroupOwner)
}

func FindMapById(id string) models.MyMap {
	var newAreaMap models.MyMap
	err := db.Select(&newAreaMap, "SELECT * FROM area_map WHERE id=?", id)

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return newAreaMap
}

func FindAllMaps() []models.MyMap {
	areaMaps := []models.MyMap{}
	err := db.Select(&areaMaps, "SELECT * FROM area_map")

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return areaMaps
}