package main

import (
	"database/sql"
	"fmt"

	menuoptions "golang_sqlite/menu_options"
	"golang_sqlite/utils"
	"log"
	"os"

	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect to database
	db, err := sql.Open("sqlite3", "./nombres.db")
	utils.CheckErr(err)
	// defer close
	defer db.Close()

	menu := wmenu.NewMenu("Que desea hacer?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Agregar una persona", 0, true, nil)
	menu.Option("Buscar una persona", 1, false, nil)
	menu.Option("Actualisar informacion de una persona", 2, false, nil)
	menu.Option("Borrar una persona por id", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:

		menuoptions.AddPerson(db)

	case 1:
		menuoptions.SearchPerson(db)

	case 2:
		menuoptions.UpdatePerson(db)
	case 3:
		menuoptions.DeletePerson(db)
	case 4:
		fmt.Println("Adios!")
		os.Exit(3)
	}
}
