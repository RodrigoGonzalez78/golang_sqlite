package databaseOptions

import (
	"database/sql"
	"fmt"
	"golang_sqlite/models"
	"golang_sqlite/utils"
	"log"
)

func AddPerson(db *sql.DB, newPerson models.Person) {

	stmt, _ := db.Prepare("INSERT INTO people (id, first_name, last_name, email, ip_address) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec(nil, newPerson.First_name, newPerson.Last_name, newPerson.Email, newPerson.Ip_address)
	defer stmt.Close()

	fmt.Printf("Agregado %v %v \n", newPerson.First_name, newPerson.Last_name)
}

func SearchForPerson(db *sql.DB, searchString string) []models.Person {

	rows, _ := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")

	defer rows.Close()

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]models.Person, 0)

	for rows.Next() {
		ourPerson := models.Person{}
		err = rows.Scan(&ourPerson.Id, &ourPerson.First_name, &ourPerson.Last_name, &ourPerson.Email, &ourPerson.Ip_address)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people
}

func GetPersonById(db *sql.DB, ourID string) models.Person {

	rows, _ := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourPerson := models.Person{}

	for rows.Next() {
		rows.Scan(&ourPerson.Id, &ourPerson.First_name, &ourPerson.Last_name, &ourPerson.Email, &ourPerson.Ip_address)
	}

	return ourPerson
}

func UpdatePerson(db *sql.DB, ourPerson models.Person) int64 {

	stmt, err := db.Prepare("UPDATE people set first_name = ?, last_name = ?, email = ?, ip_address = ? where id = ?")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(ourPerson.First_name, ourPerson.Last_name, ourPerson.Email, ourPerson.Ip_address, ourPerson.Id)
	utils.CheckErr(err)

	affected, err := res.RowsAffected()
	utils.CheckErr(err)

	return affected
}

func DeletePerson(db *sql.DB, idToDelete string) int64 {

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	utils.CheckErr(err)

	affected, err := res.RowsAffected()
	utils.CheckErr(err)

	return affected
}
