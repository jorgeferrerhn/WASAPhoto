package database

import (
	"fmt"
	"log"
)

var (
	id         uint64
	nameSearch string
)

func (db *appdbimpl) DoLogin(u User) (User, error) {

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id, name from users where name=?`, u.Name)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &nameSearch)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", nameSearch)
	if nameSearch == "" { //user has not been created before
		res, err := db.c.Exec(`INSERT INTO users (id, name,profilepic,followers,banned,photos) VALUES (NULL, ?,?,?,?,?)`,
			u.Name, 0, "[]", "[]", "[]")
		if err != nil {
			return u, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return u, err
		}

		u.ID = uint64(lastInsertID)
		u.ProfilePic = 0
		u.Followers = "[]"
		u.Banned = "[]"
		u.Photos = "[]"
	} else {
		u.ID = id
	}

	return u, nil
}
