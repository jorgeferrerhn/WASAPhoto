package database

import (
	"log"
)

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {

	var photoId uint64
	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id from photos where path=? and userid=?`, p.Path, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&photoId)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if photoId == 0 { //photo has not been uploaded before
		res, err := db.c.Exec(`INSERT INTO photos (id,userid,path,likes,comments,date) VALUES (NULL,?,?,?,?,?)`,
			p.UserId, p.Path, "[]", "[]", p.Date)
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
