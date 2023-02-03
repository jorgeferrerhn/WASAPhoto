package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) LikePhoto(p Photo, u User) (Photo, User, error) {
	var castPhotos []Photo

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

	var liked bool

	// We check that the photo hasn't been liked before
	liked = strings.Contains(p.Likes, fmt.Sprint(u.ID))

	if !liked { // Chapucero: a ver si puedo cambiarlo
		var add string
		new_list := p.Likes[0 : len(p.Likes)-1]
		if p.Likes == "[]" {
			add = fmt.Sprint(u.ID) + "]"
		} else {
			add = "," + fmt.Sprint(u.ID) + "]"
		}
		new_list += add
		p.Likes = new_list
	}

	// We update the user's photos and the photos' stream

	// Here we update the information of the photo on "raw format" { 1 Content ...} --> json.Unmarshal
	in2 := []byte(u.Photos)
	err = json.Unmarshal(in2, &castPhotos)
	if err != nil {
		return p, u, err
	}

	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { // this is the one who gets commented
			castPhotos[i].Likes = p.Likes
		}
	}

	u.ID = p.UserId // this is important: the id of the user we need to update is not the one who likes, but the one who gets the like on the photo

	// Now, we have to store castPhotos as {"ID": 1, "Content": ...} --> json.Marshal
	savePhotos, err := json.Marshal(castPhotos)
	if err != nil {
		return p, u, err
	}
	u.Photos = string(savePhotos)

	// SQL Statements
	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return p, u, err
	}

	res, err = db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
		p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
	if err != nil {
		return p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	return p, u, nil

}
