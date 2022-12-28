package database

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	id1        uint64
	name       string
	profilepic uint64
	followers  string
)

func (db *appdbimpl) GetUserProfile(id int) ([]byte, error) {

	fmt.Println("AQUI")

	rows, err := db.c.Query(`select id, name,profilepic,followers from users where id=?`, id) //Here followers will be a string, then casted to string array

	defer rows.Close()

	var u User

	for rows.Next() {
		err := rows.Scan(&id1, &name, &profilepic, followers)

		if err != nil {
			log.Fatal(err)
		}
		log.Println("this: ", id1, name, profilepic, followers)

		//cast to json
		u.ID = id1
		u.Name = name
		u.ProfilePic = profilepic
		u.Followers = followers

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(u)

	return json, err
}
