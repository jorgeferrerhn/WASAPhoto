package database

import (
	"errors"
)

func (db *appdbimpl) GetUserProfile(u User, token int) (User, error) {
	var tokName string

	// First, we have to search if the token ID exists
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, token)

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tokName)
		if err != nil {
			return u, err
		}
	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if tokName == "" {
		return u, errors.New("Token user doesn't exist!")
	}

	rows, err = db.c.Query(`select name,profilepic,followers,banned, photos from users where id=?`, u.ID) // Here followers will be a string, then cast to string array

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
