package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"

	"hello/tantan/controller"
	"log"
	"net/http"
)

func init() {

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome TanTan!\n")
}

func main() {

	r := mux.NewRouter()
	// A route with a route variable:
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users", controller.AddUser).Methods("POST")
	r.HandleFunc("/users/{user_id}/relationships", controller.GetRelationship).Methods("GET")
	r.HandleFunc("/users/{user_id}/relationships/{other_user_id}", controller.PutRelationship).Methods("PUT")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
