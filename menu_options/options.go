package menuoptions

import (
	"bufio"
	"database/sql"
	"fmt"
	databaseoptions "golang_sqlite/database_options"
	"golang_sqlite/models"
	"os"
	"strings"
)

func AddPerson(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese primer nombre: ")
	firstName, _ := reader.ReadString('\n')
	if firstName != "\n" {
		firstName = strings.TrimSuffix(firstName, "\n")
	}

	fmt.Print("Ingrese apellido: ")
	lastName, _ := reader.ReadString('\n')
	if lastName != "\n" {
		lastName = strings.TrimSuffix(lastName, "\n")
	}

	fmt.Print("Ingrese email: ")
	email, _ := reader.ReadString('\n')
	if email != "\n" {
		email = strings.TrimSuffix(email, "\n")
	}

	fmt.Print("Ingrese ip: ")
	ipAddress, _ := reader.ReadString('\n')
	if ipAddress != "\n" {
		ipAddress = strings.TrimSuffix(ipAddress, "\n")
	}

	newPerson := models.Person{
		First_name: firstName,
		Last_name:  lastName,
		Email:      email,
		Ip_address: ipAddress,
	}

	databaseoptions.AddPerson(db, newPerson)
}

func SearchPerson(db *sql.DB) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre a buscar: ")
	searchString, _ := reader.ReadString('\n')
	searchString = strings.TrimSuffix(searchString, "\n")
	people := databaseoptions.SearchForPerson(db, searchString)

	fmt.Printf("Found %v resultados\n", len(people))

	for _, ourPerson := range people {
		fmt.Printf("\n----\nNombre: %s\nApellido: %s\nEmail: %s\nIP: %s\n", ourPerson.First_name, ourPerson.Last_name, ourPerson.Email, ourPerson.Ip_address)
	}

}

func UpdatePerson(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese id a actualisar: ")
	updateid, _ := reader.ReadString('\n')

	currentPerson := databaseoptions.GetPersonById(db, updateid)

	fmt.Printf("Nombre (Actual %s):", currentPerson.First_name)
	firstName, _ := reader.ReadString('\n')
	if firstName != "\n" {
		currentPerson.First_name = strings.TrimSuffix(firstName, "\n")
	}

	fmt.Printf("Apellido (Actual %s):", currentPerson.Last_name)
	lastName, _ := reader.ReadString('\n')
	if lastName != "\n" {
		currentPerson.Last_name = strings.TrimSuffix(lastName, "\n")
	}

	fmt.Printf("Email (Actual %s):", currentPerson.Email)
	email, _ := reader.ReadString('\n')
	if email != "\n" {
		currentPerson.Email = strings.TrimSuffix(email, "\n")
	}

	fmt.Printf("IP (Actual %s):", currentPerson.Ip_address)
	ipAddress, _ := reader.ReadString('\n')
	if ipAddress != "\n" {
		currentPerson.Ip_address = strings.TrimSuffix(ipAddress, "\n")
	}

	affected := databaseoptions.UpdatePerson(db, currentPerson)

	if affected == 1 {
		fmt.Println("Una fila afectada")
	}
}

func DeletePerson(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese id a eleiminar : ")
	searchString, _ := reader.ReadString('\n')

	idToDelete := strings.TrimSuffix(searchString, "\n")

	affected := databaseoptions.DeletePerson(db, idToDelete)

	if affected == 1 {
		fmt.Println("Borrado de la base de datos!")
	}

}
