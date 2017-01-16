package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/negroni"

	"github.com/noliva/api-maps/controllers"
	"github.com/noliva/api-maps/database"
)

func main() {
	//Initialize database
	createTables(database.Connector())

	router := mux.NewRouter()

	router.HandleFunc("/maps", controllers.GetMaps).Methods("GET")
	router.HandleFunc("/maps", controllers.CreateMap).Methods("POST")
	router.HandleFunc("/maps/{id}", controllers.GetMap).Methods("GET")
	router.HandleFunc("/maps/{id}", controllers.DeleteMap).Methods("DELETE")

	middlewareHandler := negroni.New()
	middlewareHandler.UseFunc(middlewareFunction)
	middlewareHandler.UseHandler(router)

	port := ":12345"
	log.Printf("Serving at port %s", port)
	log.Fatal(http.ListenAndServe(port, middlewareHandler))
}

func middlewareFunction(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Requested: ", r.URL, r.Method, r.Host)
	rw.Header().Set("Content-type", "application/json")
	next(rw, r)
}

func createTables(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS map (
	  id int(11) NOT NULL AUTO_INCREMENT,
	  name varchar(255) NOT NULL,
	  description varchar(255) NOT NULL,
	  boundaries varchar(255) NOT NULL,
	  group_owner varchar(255) NOT NULL,
	  created_at datetime NOT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;`

	schema2 := `
	CREATE TABLE IF NOT EXISTS address (
	  id int(11) NOT NULL AUTO_INCREMENT,
	  map_id int(11) NOT NULL,
	  building_number varchar(20) DEFAULT NULL,
	  building varchar(100) DEFAULT NULL,
	  street_number varchar(20) DEFAULT NULL,
	  street varchar(100) NOT NULL,
	  post_code varchar(8) NOT NULL,
	  created_at datetime NOT NULL,
	  group_owner varchar(255) NOT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;`

	_, err := db.Exec(schema)

	if err != nil {
		log.Fatalln("Error creating tables: ", err)
	}
	_, err2 := db.Exec(schema2)

	if err2 != nil {
		log.Fatalln("Error creating tables: ", err2)
	}
}
