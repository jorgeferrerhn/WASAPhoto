package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) UnlikePhoto(p Photo, u User) (Photo, User, error) {
	var castPhotos []Photo

	// search for the user
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {

		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err2 := rows.Scan(&u.Name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err2 != nil {
			return p, u, err2
		}
	}

	if u.Name == "" {
		// el usuario no existía
		return p, u, errors.New("User not found")
	}
	err3 := rows.Err()
	if err3 != nil {
		return p, u, err3
	}

	// search the photo
	rows2, err4 := db.c.Query(`select userId,path,likes,comments,date from photos where id=?`, p.ID)

	if err4 != nil {
		return p, u, err4
	}

	defer rows2.Close()

	for rows2.Next() {

		err5 := rows2.Scan(&p.UserId, &p.Path, &p.Likes, &p.Comments, &p.Date)

		if err5 != nil {

			return p, u, err5
		}

	}

	err6 := rows2.Err()
	if err6 != nil {
		return p, u, err6
	}

	if p.Path == "" {
		// el usuario no existía
		return p, u, errors.New("Photo not found")
	}

	liked := strings.Contains(p.Likes, fmt.Sprint(u.ID))

	if !liked {
		return p, u, errors.New("Photo wasn't previously liked!")
	}

	newList := "["
	counter := 1
	// Updating the followers' list
	for i := 1; i < len(p.Likes)-1; i++ { // Chapuza: esto hay que cambiarlo
		c := string(p.Likes[i]) // rune to string
		if c == "," {
			number := p.Likes[counter:i] // takes up to that position
			if number != fmt.Sprint(u.ID) {
				newList += number + ","
			}
		}
	}
	newList = newList[:len(newList)-1] + "]" // extract the last comma and add the last bracket
	if newList == "]" {
		// It was empty
		newList = "[]"
	}

	p.Likes = newList // update the likes list

	in := []byte(u.Photos)

	// Here we update the information of the photo on "raw format" { 1 Content ...} --> json.Unmarshal
	err7 := json.Unmarshal(in, &castPhotos)
	if err7 != nil {
		return p, u, err7
	}

	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { //this is the one who gets commented
			castPhotos[i].Likes = p.Likes
		}
	}
	savePhotos, err8 := json.Marshal(castPhotos)
	if err8 != nil {
		return p, u, err8
	}
	u.Photos = string(savePhotos)
	u.ID = p.UserId // this is important: the id of the user we need to update is not the one who likes, but the one who gets the like on the photo

	res, err9 := db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err9 != nil {
		return p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	res2, err10 := db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
		p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
	if err10 != nil {
		return p, u, errors.New("Error in: " + fmt.Sprint(res2))
	}

	return p, u, nil

}
