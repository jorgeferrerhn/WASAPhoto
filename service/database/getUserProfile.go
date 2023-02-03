package database

import (
	"errors"
)

func (db *appdbimpl) GetUserProfile(u User, token int) (User, error) {
	var id int
	var name, tokName string

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

	rows, err = db.c.Query(`select id, name,profilepic,followers,banned, photos from users where id=?`, u.ID) // Here followers will be a string, then cast to string array

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err != nil {
			return u, err
		}

	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if name == "" || id == 0 {
		return u, errors.New("User not found")
	}

	u.Name = name
	return u, nil
}
