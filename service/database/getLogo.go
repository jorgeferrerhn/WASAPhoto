package database

import (
	"log"
)

var (
	searchedLogo uint64
)

func (db *appdbimpl) GetLogo(id int) (uint64, error) {

	rows, err := db.c.Query(`select profilepic from users where id=?`, id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&searchedLogo)

		if err != nil {
			log.Fatal(err)
		}
		//log.Println("this: ", id1, name, profilepic, followers, photos)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return searchedLogo, err
}
