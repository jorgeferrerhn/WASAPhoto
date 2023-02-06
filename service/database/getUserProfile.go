package database

import (
	"errors"
)

func (db *appdbimpl) GetUserProfile(u User) (User, error) {

	rows, err := db.c.Query(`select name,profilepic,followers,banned, photos from users where id=?`, u.ID) // Here followers will be a string, then cast to string array

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err != nil {
			return u, err
		}

	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if u.Name == "" || u.ID == 0 {
		return u, errors.New("User not found")
	}

	return u, nil
}
