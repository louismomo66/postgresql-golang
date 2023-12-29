package controllers

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go_postgtresql_pgx/models"
	"log"
	"time"
)

func GetAll(dbpool *pgxpool.Pool) ([]models.Person, error) {
	var persons []models.Person
	rows, err := dbpool.Query(context.Background(), "SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Person
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func CreateNew(dbpool *pgxpool.Pool, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
	var p models.Person
	err := dbpool.QueryRow(context.Background(), "INSERT INTO person (first_name,last_name,date_of_birth)VALUES($1,$2,$3) RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth.Format("2006-01-02")).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
	return p, err

}

func DeletePerson(dbpool *pgxpool.Pool, id int) error {
	_, err := dbpool.Exec(context.Background(), "DELETE FROM person WHERE id= $1", id)
	return err
}

func UpdatePerson(dbpool *pgxpool.Pool, id int, firstname, lastname string, dateofbirth time.Time) (models.Person, error) {
	var p models.Person
	err := dbpool.QueryRow(context.Background(), "UPDATE person SET first_name=$1,last_name=$2,date_of_birth=$3 WHERE id=$4 RETURNING id,first_name,last_name,date_of_birth", firstname, lastname, dateofbirth, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
	return p, err
}
func GetOne(dbpool *pgxpool.Pool, id int) models.Person {
	var p models.Person
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM person WHERE id = $1", id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth)
	if err != nil {
		log.Fatalf("Person with Id: %d doesnot exist %v\n ", id, err)
	}
	return p
}
