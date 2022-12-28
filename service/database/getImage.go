package database

import (
	"log"
)

func (db *appdbimpl) GetImage(userId int, imageId int) (uint64, error) {

	var searchedPhotos string

	rows, err := db.c.Query(`select photos from users where id=?`, userId)

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

	return 0, err
}
