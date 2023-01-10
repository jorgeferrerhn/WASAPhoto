package database

import (
	"fmt"
	"log"
)

func (db *appdbimpl) SetMyUserName(id int, name string) (int, error) {

	fmt.Println(id)
	fmt.Println(name)

	var tname string
	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select name from users where id=?`, id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&tname)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if tname == "" { //invalid user id
		return -1, nil
	}

	//falta actualizar la tabla de usuarios

	return 1, nil
}
