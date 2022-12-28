package database

import (
	"fmt"
	"log"
)

func (db *appdbimpl) GetImage(userId int, imageId int) (byte, error) {

	var searchedPhotos string
	var img byte

	rows, err := db.c.Query(`select photos from users where id=?`, userId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&searchedPhotos)

		fmt.Println(err)

		if err != nil {
			log.Fatal(err)
		}
		//log.Println("this: ", id1, name, profilepic, followers, photos)

		//at this point, we access the imageID photo and return it
		img = searchedPhotos[imageId] //aquí está el error: tenemos que probar con un ejemplo que funcione

	}
	return img, err
}
