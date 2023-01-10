package database

import (
	"fmt"
	"log"
	"time"
)

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {

	var photoId uint64
	var userName string

	//search for the user
	rows2, err := db.c.Query(`select name from users where id=?`, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&userName)

		fmt.Println("Previous user name: ", userName)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	//search for the photo id (check if it existed)
	rows, err := db.c.Query(`select id from photos where path=? and userid=?`, p.Path, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&photoId)

		fmt.Println("Photo id: ", photoId)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if userName != "" && photoId == 0 { //photo has not been uploaded before

		p.Likes = "[]"
		p.Comments = "[]"
		p.Date = time.Now()

		res, err := db.c.Exec(`INSERT INTO photos (id,userid,path,likes,comments,date) VALUES (NULL,?,?,?,?,?)`,
			p.UserId, p.Path, p.Likes, p.Comments, p.Date)
		if err != nil {
			return p, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return p, err
		}

		p.ID = uint64(lastInsertID)

	} else {
		p.ID = photoId
	}

	return p, nil
}
