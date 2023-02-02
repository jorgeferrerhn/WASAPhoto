package database

func (db *appdbimpl) DoLogin(u User) (User, error) {

	var id int
	var nameSearch string

	// first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id, name,profilepic,followers,banned,photos from users where name=?`, u.Name)

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&id, &nameSearch, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)
		if err != nil {
			return u, err
		}
	}

	err = rows.Err()
	if err != nil {
		return u, err
	}

	if nameSearch == "" || id == 0 { // this user has not been created before

		u.ProfilePic = 0
		u.Followers = "[]"
		u.Banned = "[]"
		u.Photos = "[]"

		res, e := db.c.Exec(`INSERT INTO users (id, name,profilepic,followers,banned,photos) VALUES (NULL, ?,?,?,?,?)`,
			u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos)

		if e != nil {
			return u, e
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return u, err
		}

		u.ID = int(lastInsertID) // gets the ID

	} else { // This user has been created before
		u.ID = id

		rows, err = db.c.Query(`select followers,profilepic,banned,photos from users where id=?`, u.ID)

		if err != nil {
			return u, err
		}

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&u.Followers, &u.ProfilePic, &u.Banned, &u.Photos)
			if err != nil {
				return u, err
			}
		}

		err = rows.Err()
		if err != nil {
			return u, err
		}

	}

	return u, nil
}
