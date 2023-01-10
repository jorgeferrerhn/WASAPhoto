package database

import (
	"log"
	"time"
)

func (db *appdbimpl) GetImage(p Photo) (Photo, error) {

	var userId int
	var path, likes, comments string
	var date time.Time

	rows, err := db.c.Query(`select userid,path,likes,comments,date from photos where id=?`, p.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userId, &path, &likes, &comments, &date)

		//we update the information of our struct
		p.UserId = userId
		p.Path = path
		p.Likes = likes
		p.Comments = comments
		p.Date = date

		if err != nil {
			log.Fatal(err)
		}
		log.Println("this: ", p.UserId, p.Path, p.Likes, p.Comments, p.Date)

	}
	return p, err
}
