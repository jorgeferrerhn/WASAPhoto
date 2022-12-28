package database

import (
	"encoding/json"
	"log"
)

var (
	id1  uint64
	name string
)

func (db *appdbimpl) GetUserProfile(id int) ([]byte, error) {

	rows, err := db.c.Query(`select id, name from users where id=?`, id)

	defer rows.Close()

	var u User

	for rows.Next() {
		err := rows.Scan(&id1, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("this: ", id1, name)

		//cast to json
		u.ID = id1
		u.Name = name

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(u)

	return json, err
}
