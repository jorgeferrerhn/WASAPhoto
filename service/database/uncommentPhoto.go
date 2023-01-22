package database

import (
	"fmt"
)

func (db *appdbimpl) UncommentPhoto(c Comment) (int, error) {

	var commentId int
	// var userId int

	/*
		//lastly, we check that the user existed
			rows3, err := db.c.Query(`select id from users where id=?`, c.UserId)

			if err != nil {
				return 0,err
			}

			defer rows3.Close()

			for rows3.Next() {
				err := rows.Scan(&userId)

				if err != nil {


					return 0,err
				}


			}

			err = rows3.Err()
			if err != nil {
				return 0,err
			}
	*/

	//first we search the comment. It should have a unique commentId, so we'll search for it
	rows, err := db.c.Query(`select commentid from comments where commentid=?`, c.ID)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&commentId)

		if err != nil {

			return 0, err
		}
	}

	err = rows.Err()
	if err != nil {
		return 0, err
	}

	//then we search the photo id. If it doesn't exist, we cannot comment on the photo

	rows2, err := db.c.Query(`select * from photos where id=?`, c.PhotoId)

	if err != nil {
		return 0, err
	}

	defer rows2.Close()

	err = rows2.Err()
	if err != nil {
		return 0, err
	}

	//esto es lo que tenemos que cambiar

	res, err := db.c.Exec(`DELETE FROM comments WHERE commentid=?`,
		c.ID)
	if err != nil {
		return 1, err
	}
	fmt.Println(res)

	return 1, nil

}
