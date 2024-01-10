package controllers

import (
	"encoding/json"
	"fmt"
	"go_postgtresql_pgx/database"
	"go_postgtresql_pgx/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// func GetAll(dbpool *pgxpool.Pool) ([]models.Person, error) {
// 	var persons []models.Person
// 	rows, err := dbpool.Query(context.Background(), "SELECT * FROM person")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var p models.Person
// 		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
// 		if err != nil {
// 			return nil, err
// 		}
// 		persons = append(persons, p)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

//		return persons, nil
//	}
//
//	func GetAll(db *gorm.DB) ([]models.Person, error) {
//		var persons []models.Person
//		results := db.Find(&persons)
//		if results.Error != nil {
//			return nil, results.Error
//		}
//		return persons, nil
//	}
func GetAll(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	results := database.DB.Find(&persons)
	if results.Error != nil {
		http.Error(w, results.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(persons)
}

// func CreateNew(dbpool *pgxpool.Pool, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
// 	var p models.Person
// 	err := dbpool.QueryRow(context.Background(), "INSERT INTO person (first_name,last_name,date_of_birth)VALUES($1,$2,$3) RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth.Format("2006-01-02")).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
// 	return p, err

// }
func CreatNew(db *gorm.DB, person models.Person) {
	result := db.Create(&person)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

//	func DeletePerson(dbpool *pgxpool.Pool, id int) error {
//		_, err := dbpool.Exec(context.Background(), "DELETE FROM person WHERE id= $1", id)
//		return err
//	}
func DeletePerson(db *gorm.DB, id int) {
	var p models.Person
	result := db.Delete(&p, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

//	func UpdatePerson(dbpool *pgxpool.Pool, id int, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
//		var p models.Person
//		err := dbpool.QueryRow(context.Background(), "UPDATE person SET first_name=$1,last_name=$2,date_of_birth=$3 WHERE id=$4 RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
//		return p, err
//	}
func UpdatePerson(db *gorm.DB, id int, updateData map[string]interface{}) {
	var p models.Person
	result := db.Model(&p).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

// func GetOne(dbpool *pgxpool.Pool, id int) models.Person {
// 	var p models.Person
// 	err := dbpool.QueryRow(context.Background(), "SELECT * FROM person WHERE id = $1", id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
// 	if err != nil {
// 		log.Fatalf("Person with Id: %d doesnot exist %v\n ", id, err)
// 	}
// 	return p
// }

// get one person
//
//	func GetOne(db *gorm.DB, id int) models.Person {
//		var p models.Person
//		result := db.First(&p, id)
//		if result.Error != nil {
//			log.Fatalf("Non existent %v\n", result.Error)
//		}
//		return p
//	}
func GetOne(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Person
	if err := database.DB.Where("id = ?", id).First(&p).Error; err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	json.NewEncoder(w).Encode(p)
}
