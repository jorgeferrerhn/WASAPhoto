package database

import (
	"errors"
)

func (db *appdbimpl) GetUserProfile(u User) (User, error) {
	var id int
	var name string
	rows, err := db.c.Query(`select id, name,profilepic,followers,banned, photos from users where id=?`, u.ID) // Here followers will be a string, then cast to string array

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

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
