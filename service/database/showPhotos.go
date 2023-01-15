package database

import (
	"log"
	"time"
)

func (db *appdbimpl) ShowPhotos(ps Photos) (Photos, error) {
	var id, userid int
	var path, likes, comments string
	var date time.Time

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id, userid,path,likes,comments,date from photos`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &userid, &path, &likes, &comments, &date)
		if err != nil {
			log.Fatal(err)
		}

		//we create and append a photo for each row
		var p Photo
		p.ID = id
		p.UserId = userid
		p.Likes = likes
		p.Comments = comments
		p.Date = date

		//we add the element to the photos list
		ps.Photos = append(ps.Photos, p)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ps, nil
}
