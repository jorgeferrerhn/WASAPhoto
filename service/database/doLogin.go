package database

import "fmt"

func (db *appdbimpl) CreateUser(u User) (User, error) {
	fmt.Println("Estamos a punto de introducir al usuario: ", u.Name)

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
