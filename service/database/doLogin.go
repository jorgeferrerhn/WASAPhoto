package database

func (db *appdbimpl) DoLogin(u User) (User, error) {

	var (
		id, profilepic                        int
		nameSearch, photos, banned, followers string
	)

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id, name from users where name=?`, u.Name)

	castError(err)

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &nameSearch)
		castError(err)
	}

	err = rows.Err()
	castError(err)

	if nameSearch == "" || nameSearch != u.Name { //this user has not been created before

		u.ProfilePic = 0
		u.Followers = "[]"
		u.Banned = "[]"
		u.Photos = "{}"

		res, err := db.c.Exec(`INSERT INTO users (id, name,profilepic,followers,banned,photos) VALUES (NULL, ?,?,?,?,?)`,
			u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos)

		if err != nil {
			return u, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return u, err
		}

		u.ID = int(lastInsertID) //gets the ID

	} else {
		u.ID = id
		rows, err := db.c.Query(`select followers,profilepic,banned,photos from users where id=?`, u.ID)

		castError(err)

		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(&followers, &profilepic, &banned, &photos)
			castError(err)
		}

		err = rows.Err()
		castError(err)

		u.ProfilePic = profilepic
		u.Followers = followers
		u.Banned = banned
		u.Photos = photos

	}

	return u, nil
}
