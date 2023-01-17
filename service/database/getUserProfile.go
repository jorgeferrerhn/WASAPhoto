package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfile(u User) (User, error) {
	var id, profilepic int
	var name, followers, banned, photos string
	rows, err := db.c.Query(`select id, name,profilepic,followers,banned, photos from users where id=?`, u.ID) //Here followers will be a string, then casted to string array

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &profilepic, &followers, &banned, &photos)

		if err != nil {
			log.Fatal(err)
		}

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if name == "" {
		return u, errors.New("User not found")
	}

	u.ID = id
	u.Name = name
	u.ProfilePic = profilepic
	u.Followers = followers
	u.Banned = banned
	u.Photos = photos

	return u, nil
}
