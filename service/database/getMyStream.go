package database

import (
	"errors"
)

func (db *appdbimpl) GetMyStream(u User) (User, error) {

	var id int
	var username string

	rows, err := db.c.Query(`select id,name,photos from users where id=?`, u.ID) // Here photos will be a string, then casted to json

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username, &u.Photos)

		if err != nil {
			return u, err
		}

	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if username == "" || id == 0 {
		return u, errors.New("User not found")
	}

	return u, err
}
