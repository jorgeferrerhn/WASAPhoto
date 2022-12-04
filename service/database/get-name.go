package database

func (db *appdbimpl) doLogin(u User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO fountains (name) VALUES (?)`, //falta
		u.Name)
	if err != nil {
		return u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	f.ID = uint64(lastInsertID)
	return f, nil
} //comprobar aquí
