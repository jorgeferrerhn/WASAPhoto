package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

func (db *appdbimpl) DeletePhoto(p Photo, u User) (Photo, User, error) {

	var castPhotos, newPhotos []Photo

	// search for the user
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {

		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&u.Name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err != nil {
			return p, u, err
		}
	}

	if u.Name == "" {
		// el usuario no existía
		return p, u, errors.New("User not found")
	}
	err = rows.Err()
	if err != nil {
		return p, u, err
	}

	// search the photo
	rows2, err2 := db.c.Query(`select userId,path,likes,comments,date from photos where id=?`, p.ID)

	if err2 != nil {
		return p, u, err2
	}

	defer rows2.Close()

	for rows2.Next() {

		err = rows2.Scan(&p.UserId, &p.Path, &p.Likes, &p.Comments, &p.Date)

		if err != nil {

			return p, u, err
		}

	}

	err = rows2.Err()
	if err != nil {
		return p, u, err
	}

	if p.Path == "" {
		// el usuario no existía
		return p, u, errors.New("Photo not found")
	}

	in := []byte(u.Photos)

	// Here we update the information of the photo on "raw format" { 1 Content ...} --> json.Unmarshal
	err = json.Unmarshal(in, &castPhotos)
	if err != nil {
		return p, u, err
	}

	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID != p.ID { // we add everything except the comments
			newPhotos = append(newPhotos, castPhotos[i])
		}
	}
	savePhotos, err := json.Marshal(newPhotos)
	u.Photos = string(savePhotos)

	// SQL Statements
	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return p, u, err
	}

	res, err = db.c.Exec(`DELETE FROM photos WHERE id=?`, p.ID)
	if err != nil {
		return p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	return p, u, nil

}
