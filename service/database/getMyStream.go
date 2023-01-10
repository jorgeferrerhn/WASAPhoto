package database

import (
	"fmt"
	"log"
)

var (
	searchedPhotos string
)

func (db *appdbimpl) GetMyStream(id int) (string, error) {

	rows, err := db.c.Query(`select photos from users where id=?`, id) //Here photos will be a string, then casted to json

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&searchedPhotos)

		if err != nil {
			log.Fatal(err)
		}
		//log.Println("this: ", id1, name, profilepic, followers, photos)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Searched photos: ", searchedPhotos)

	return searchedPhotos, err
}
