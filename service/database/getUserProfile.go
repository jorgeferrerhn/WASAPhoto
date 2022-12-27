package database

import (
	"database/sql"
	"log"
)

var (
	id   int
	name string
)

func (db *appdbimpl) GetUserProfile(id int) (sql.Rows, error) {

	rows, err := db.c.Query("select id, name from users")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return *rows, err
}
