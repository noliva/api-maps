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

	middlewareHandler := negroni.New()
	middlewareHandler.UseFunc(middlewareFunction)
	middlewareHandler.UseHandler(router)

	port := ":12345"
	log.Printf("Serving at port %s", port)
	log.Fatal(http.ListenAndServe(port, middlewareHandler))
}

func middlewareFunction(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("middleware", r.URL)

	rw.Header().Set("Content-type", "application/json")
	next(rw, r)
}

func createTables(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS area_map (
	id varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	boundaries varchar(255) NOT NULL,
	group_owner varchar(255) NOT NULL,
	PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := db.Exec(schema)

	if err != nil {
		log.Fatalln("Error creating tables: ", err)
	}
}
