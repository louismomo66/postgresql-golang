package controllers

import (
	"encoding/json"
	"go_postgtresql_pgx/models"
	"go_postgtresql_pgx/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PersonController struct {
	DB *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{DB: db}
}
func (pc *PersonController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var persons []models.Person
	results := pc.DB.Find(&persons)
	if results.Error != nil {
		http.Error(w, results.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(persons)
}
func (pc *PersonController) GetOne(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Person
	if err := pc.DB.Where("id = ?", id).First(&p).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "person not found")
		return
	}
	json.NewEncoder(w).Encode(p)
}
func (pc *PersonController) CreatNew(w http.ResponseWriter, r *http.Request) {
	var p models.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := pc.DB.Create(&p)
	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
func (pc *PersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var p models.Person
	if err := pc.DB.Where("id = ?", id).First(&p).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Person not there")
		return
	}

	pc.DB.Delete(&p)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(p)
}
func (pc *PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var p models.Person
	if err := pc.DB.Where("id = ?", id).First(&p).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Person not there")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var newp models.Person
	newp.FirstName = p.FirstName
	newp.LastName = p.LastName
	newp.DateOfBirth = p.DateOfBirth

	pc.DB.Save(&newp)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(newp)
}

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

// func CreateNew(dbpool *pgxpool.Pool, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
// 	var p models.Person
// 	err := dbpool.QueryRow(context.Background(), "INSERT INTO person (first_name,last_name,date_of_birth)VALUES($1,$2,$3) RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth.Format("2006-01-02")).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
// 	return p, err

// }

//	func DeletePerson(dbpool *pgxpool.Pool, id int) error {
//		_, err := dbpool.Exec(context.Background(), "DELETE FROM person WHERE id= $1", id)
//		return err
//	}

//	func UpdatePerson(dbpool *pgxpool.Pool, id int, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
//		var p models.Person
//		err := dbpool.QueryRow(context.Background(), "UPDATE person SET first_name=$1,last_name=$2,date_of_birth=$3 WHERE id=$4 RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
//		return p, err
//	}

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
