package database

import (
	"fmt"
	"log"
)

func (db *appdbimpl) DeletePhoto(p Photo) (int, error) {

	var photoId uint64
	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id from photos where id=?`, p.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&photoId)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`,
		photoId)

	fmt.Println(res)
	if err != nil {
		return -1, err
	}

	return 1, nil
}
