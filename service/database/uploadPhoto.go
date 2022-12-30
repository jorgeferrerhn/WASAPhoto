package database

import (
	"log"
	"time"
)

type Photo struct {
	Path     string
	PhotoID  uint64
	Likes    uint64
	Comments string
	Date     time.Time
}

func (db *appdbimpl) UploadPhoto(userId int, path string) (int, error) {

	var selectedPhotos string

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select photos from users where id=?`, id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&selectedPhotos)

		//create struct and cast

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
