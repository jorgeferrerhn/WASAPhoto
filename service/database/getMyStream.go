package database

import (
	"errors"
)

func (db *appdbimpl) GetMyStream(u User) (User, error) {

	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, u.ID) // Here photos will be a string, then cast to json

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err2 := rows.Scan(&u.Name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err2 != nil {
			return u, err2
		}

	}
	err3 := rows.Err()
	if err3 != nil {
		return u, err3
	}

	if u.Name == "" || u.ID == 0 {
		return u, errors.New("User not found")
	}

	return u, nil
}
