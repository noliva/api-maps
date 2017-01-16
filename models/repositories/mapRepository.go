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

func Create(m models.MyMap) {
	query := `INSERT INTO map (name, description, boundaries, group_owner, created_at) VALUES (?, ?, ?, ?, ?)`
	db.MustExec(query, m.Name, m.Description, m.Boundaries, m.GroupOwner, m.CreatedAt)
}

func Delete(m models.MyMap) {
	err := db.MustExec("DELETE FROM map WHERE id=?", m.ID)

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}
}

func FindMapById(id string) models.MyMap {
	var newAreaMap models.MyMap
	err := db.Get(&newAreaMap, "SELECT * FROM map WHERE id=?", id)

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return newAreaMap
}

func FindAllMaps() []models.MyMap {
	areaMaps := []models.MyMap{}
	err := db.Select(&areaMaps, "SELECT * FROM map")

	if err != nil {
		log.Fatalln("Error fetching: ", err)
	}

	return areaMaps
}