package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO users (id, name) VALUES (NULL, ?)`,
		u.Name)
	if err != nil {
		return u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = uint64(lastInsertID)
	return u, nil
}
