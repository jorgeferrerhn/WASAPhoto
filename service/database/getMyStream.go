package database

import "errors"

func (db *appdbimpl) GetMyStream(u User) (User, error) {

	var searchedPhotos, username string

	rows, err := db.c.Query(`select name,photos from users where id=?`, u.ID) //Here photos will be a string, then casted to json

	castError(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&username, &searchedPhotos)

		castError(err)

	}
	err = rows.Err()
	castError(err)

	if username == "" {
		return u, errors.New("User not found")
	}

	u.Photos = searchedPhotos

	return u, err
}
