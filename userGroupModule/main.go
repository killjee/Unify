package main


import (
	"userGroupModule/createUserGroup"
	"userGroupModule/getParentUserGroup"
	"userGroupModule/getAdminUserUuid"
	"userGroupModule/addUserToUserGroup"
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

func route(router *mux.Router) {
	// router.HandleFunc("/", createUserGroup.RequestHandler).Methods("GET")
	router.HandleFunc("/userGroup/create", createUserGroup.RequestHandler).Methods("POST")
	router.HandleFunc("/userGroup/parent", getParentUserGroup.RequestHandler).Methods("POST")
	router.HandleFunc("/userGroup/admin", getAdminUserUuid.RequestHandler).Methods("POST")
	router.HandleFunc("/userGroup/addUser", addUserToUserGroup.RequestHandler).Methods("POST")
}