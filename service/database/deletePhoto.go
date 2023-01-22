package database

import (
	"errors"
	"fmt"
)

func (db *appdbimpl) DeletePhoto(p Photo) (int, error) {

	var photoId int
	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id from photos where id=?`, p.ID)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&photoId)

		if err != nil {
			return 0, err
		}
	}

	err = rows.Err()
	if err != nil {
		return 0, err
	}

	res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`,
		photoId)

	if err != nil {
		return -1, errors.New("Error in " + fmt.Sprint(res))
	}

	return 1, nil
}
