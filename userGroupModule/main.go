package main


import (
	"userGroupModule/createUserGroup"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	
	port := "8080" 

	route(router)
	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

func getDb() *sql.DB {
	connStr := "user=postgres dbname=unify password=123456"
	db, err := sql.Open("postgres", connStr)
	if err != nil {	
		fmt.Println("Jhandu balm")
	}

	return db
}

func route(router *mux.Router) {
	router.HandleFunc("/", createUserGroup.RequestHandler).Methods("GET")
}