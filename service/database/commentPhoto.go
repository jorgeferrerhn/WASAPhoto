package database

import (
	"fmt"
	"log"
	"time"
)

func (db *appdbimpl) CommentPhoto(c Comment) (Comment, error) {

	var photoId uint64
	var userName string
	// var userId uint64

	//search for the user
	rows, err := db.c.Query(`select name from users where id=?`, c.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName)

		fmt.Println("Previous user name: ", userName)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//then we search the photo id. If it doesn't exist, we cannot comment on the photo

	rows2, err := db.c.Query(`select id from photos where id=?`, c.PhotoId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&photoId)

		fmt.Println("Previous photo: ", photoId)

		if err != nil {
			log.Fatal(err)
		}

	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	if userName != "" && photoId != 0 { //comment has not been uploaded before and the user and photo exist
		c.Date = time.Now()

		res, err := db.c.Exec(`INSERT INTO comments (commentid,content,photoid,userid,date) VALUES (NULL,?,?,?,?)`,
			c.ID, c.Content, c.PhotoId, c.UserId, c.Date)
		if err != nil {
			return c, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return c, err
		}

		c.ID = uint64(lastInsertID)

	}

	return c, nil

}
