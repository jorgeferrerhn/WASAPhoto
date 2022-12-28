package database

import (
	"encoding/json"
	"log"
)

var (
	id1        uint64
	name       string
	profilepic uint64
	followers  string
	photos     string
)

func (db *appdbimpl) GetUserProfile(id int) ([]byte, error) {

	rows, err := db.c.Query(`select id, name,profilepic,followers,photos from users where id=?`, id) //Here followers will be a string, then casted to string array

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var u User

	for rows.Next() {
		err := rows.Scan(&id1, &name, &profilepic, &followers, &photos)

		if err != nil {
			log.Fatal(err)
		}
		//log.Println("this: ", id1, name, profilepic, followers, photos)

		//cast to json
		u.ID = id1
		u.Name = name
		u.ProfilePic = profilepic
		u.Followers = followers
		u.Photos = photos

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(u)

	return json, err
}
