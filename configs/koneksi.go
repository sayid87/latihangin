package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Koneksi() (*sql.DB, error) {
	const userDB string = "root"
	const passDB string = ""
	const hostDB string = "localhost"
	const portDB string = "3306"
	const nameDB string = "dbgin"
	const conDB string = userDB + ":" + passDB + "@tcp(" + hostDB + ":" + portDB + ")/" + nameDB + "?parseTime=true"

	db, err := sql.Open("mysql", conDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func QueryWrite(q string) {
	db, err := Koneksi()
	if err != nil {
		log.Fatal(err.Error()) //die() klo di PHP
	}

	_, errQuery := db.Exec(q)
	if errQuery != nil {
		log.Fatal(errQuery.Error())
	}

	defer db.Close()
}

func QueryRead(q string) *sql.Rows {
	db, err := Koneksi()
	if err != nil {
		log.Fatal(err.Error()) //die() klo di PHP
	}
	result, errQuery := db.Query(q)
	defer db.Close()
	if errQuery != nil {
		log.Fatal(errQuery.Error())
		return nil
	} else {

		return result
	}

}
