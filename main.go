package main

import (
	"go_postgtresql_pgx/controllers"
	"go_postgtresql_pgx/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	database.ConnectDB()
	r.HandleFunc("/people", controllers.GetAll).Methods("GET")
	r.HandleFunc("/people/{id}", controllers.GetOne).Methods("GET")
	r.HandleFunc("/people", controllers.CreatNew).Methods("POST")
	r.HandleFunc("/people/{id}", controllers.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", controllers.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// sqlDB, err := db.DB()
// if err != nil {
// 	log.Fatalf("Failed to get the underlying sql.DB: %v", err)
// }
// defer sqlDB.Close()

// r.HandelFunc("/people/{id}", controllers.GetOne()).Methods("GET")
// r.HandelFunc("/people", controllers.CreatNew()).Methods("POST")
// r.HandelFunc("/people/{id}", controllers.DeletePerson()).Methods("DELETE")
// controllers.DeletePerson(db, 9)
// dateOfBirth := time.Date(1996, time.June, 23, 0, 0, 0, 0, time.UTC)
// updatedData := map[string]interface{}{"FirstName": "Jane", "LastName": "Doe", "DateOfBirth": dateOfBirth}
// controllers.UpdatePerson(db, 10, updatedData)
// newPerson := models.Person{FirstName: "Tom", LastName: "Dumba", DateOfBirth: time.Now()}
// controllers.CreatNew(db, newPerson)
// getting all persons
// persons, err := controllers.GetAll(db)
// if err != nil {
// 	log.Fatalf("Failed to get all error: %v", err)
// }

// for _, person := range persons {
// 	log.Printf("ID: %d, Name: %s %s, DOB: %s\n", person.ID, person.FirstName, person.LastName, person.DateOfBirth.Format("2006-01-02"))
// }

// //get one person
// person := controllers.GetOne(db, 6)
// log.Println(person)
// }

// // getting all persons
// persons, err := controllers.GetAll(dbpool)
// if err != nil {
// 	log.Printf("Error retrieving persons %v\n", err)

// }

// for _, person := range persons {
// 	log.Printf("ID: %d, Name: %s %s, DOB: %s \n", person.ID, person.FirstName, person.LastName, person.DateOfBirth.Format("2006-01-02"))
// }
// // creating a new person
// dateOfBirth := time.Date(1996, time.June, 23, 0, 0, 0, 0, time.UTC)
// // person, err := controllers.CreateNew(dbpool, "Joseph", "Damulila", dateOfBirth)
// // if err != nil {
// // 	log.Fatalf("Error creating person %v", err)
// // }
// // log.Println("Created new person:", person)
// updated, err := controllers.UpdatePerson(dbpool, 6, "Jonas", "Excellency", dateOfBirth)
// if err != nil {
// 	log.Fatalf("Failed to update %v\n", err)
// }
// log.Println("Updated: ", updated)
// individual := controllers.GetOne(dbpool, 6)
// log.Println("Individual: ", individual)
