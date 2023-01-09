package database

func (db *appdbimpl) LikePhoto(photoId int) (int, error) {

	// var userId uint64

	/*
			//lastly, we check that the user existed
				rows3, err := db.c.Query(`select id from users where id=?`, c.UserId)

				if err != nil {
					log.Fatal(err)
				}

				defer rows3.Close()

				for rows3.Next() {
					err := rows.Scan(&userId)

					if err != nil {
						fmt.Println("El error")

						log.Fatal(err)
					}

					fmt.Println("User id: ", userId)
				}

				err = rows3.Err()
				if err != nil {
					log.Fatal(err)
				}

		//first we search the comment. It should have a unique commentId, so we'll search for it
		rows, err := db.c.Query(`select commentid from comments where commentid=?`, c.ID)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(&commentId)

			if err != nil {

				log.Fatal(err)
			}
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		//then we search the photo id. If it doesn't exist, we cannot comment on the photo

		rows2, err := db.c.Query(`select * from photos where id=?`, c.PhotoId)

		if err != nil {
			log.Fatal(err)
		}

		defer rows2.Close()

		for rows2.Next() {

			fmt.Println("Photo id: ", rows2)
		}

		err = rows2.Err()
		if err != nil {
			log.Fatal(err)
		}

		if commentId == 0 { //comment has not been uploaded before
			res, err := db.c.Exec(`INSERT INTO comments (commentid,content,photoid,userid,date) VALUES (NULL,?,?,?,?)`,
				c.ID, c.Content, c.PhotoId, c.UserId, c.Date)
			if err != nil {
				return c, err
			}

			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return c, err
			}

			c.ID = uint64(lastInsertID)

		} else {
			c.ID = commentId
		}
	*/

	return 0, nil

}
