package routes

import (
	"go_postgtresql_pgx/controllers"

	"github.com/gorilla/mux"
)

func SetUpRoutes(r *mux.Router, personController *controllers.PersonController) {
	r.HandleFunc("/people", personController.GetAll).Methods("GET")
	r.HandleFunc("/people/{id}", personController.GetOne).Methods("GET")
	r.HandleFunc("/people", personController.CreatNew).Methods("POST")
	r.HandleFunc("/people/{id}", personController.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", personController.DeletePerson).Methods("DELETE")
}
