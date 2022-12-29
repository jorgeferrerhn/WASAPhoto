package database

import "fmt"

func (db *appdbimpl) UploadPhoto(path string) (int, error) {

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select id, name from users`)
	fmt.Println(rows)
	return 0, err
}
