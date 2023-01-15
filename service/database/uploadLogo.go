package database

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func (db *appdbimpl) UploadLogo(p Photo, u User) (Photo, User, error) {
	var photoId, profilePic int
	var userName, followers, banned, photos string

	//search for the user
	rows2, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&userName, &profilePic, &followers, &banned, &photos)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	//comprobamos si existía
	if userName == "" {
		//el usuario no existía
		return p, u, errors.New("User not found")
	}

	//search for the photo id (check if it existed)
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

	if userName != "" && photoId == 0 { //photo has not been uploaded before
		p.Likes = "[]"
		p.Comments = "[]"
		p.Date = time.Now()

		//We insert the new photo to the database collection
		res, err := db.c.Exec(`INSERT INTO photos (id,userid,path,likes,comments,date) VALUES (NULL,?,?,?,?,?)`,
			p.UserId, p.Path, p.Likes, p.Comments, p.Date)
		if err != nil {
			return p, u, err
		}
		fmt.Println("Followers: ", u.Followers)

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return p, u, err
		}

		p.ID = int(lastInsertID)

		//We update the user
		u.Name = userName
		u.ProfilePic = p.ID // we update the profile picture ID
		u.Followers = followers
		u.Banned = banned
		u.Photos = photos
		res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`, u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
		if err != nil {
			return p, u, err
		}

	} else {
		p.ID = photoId
	}

	return p, u, nil
}
