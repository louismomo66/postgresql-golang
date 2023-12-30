package main

import (
	"go_postgtresql_pgx/database"
	"log"
)

func main() {

	db := database.ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get the underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()

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

}
