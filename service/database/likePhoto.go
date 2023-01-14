package database

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func (db *appdbimpl) LikePhoto(p Photo) (Photo, error) {

	var likes, path, comments string
	var userId int
	var date time.Time
	var userName string

	//search for the user
	rows, err := db.c.Query(`select name from users where id=?`, p.UserId)

	if err != nil {

		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName)

		if err != nil {
			return p, err
		}
	}

	if userName == "" {
		//el usuario no existía
		return p, errors.New("User not found")
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//search the photo
	rows2, err := db.c.Query(`select userId,path,likes,comments,date from photos where id=?`, p.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&userId, &path, &likes, &comments, &date)

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(likes)
	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	if path == "" {
		//el usuario no existía
		return p, errors.New("Photo not found")
	}

	if userName != "" { //user exists

		var liked bool

		// We check that the photo hasn't been liked before
		strId := fmt.Sprint(p.UserId)
		liked = strings.ContainsAny(likes, strId)
		fmt.Println("Likes : ", likes, " of ", p.UserId, ": ", liked)

		fmt.Println("Liked: ", liked)

		if !liked {
			var add string

			new_list := likes[0 : len(likes)-1]
			strUID := fmt.Sprint(p.UserId)

			if likes == "[]" {
				add = strUID + "]"

			} else {
				add = "," + strUID + "]"

			}
			new_list += add

			p.Likes = new_list

			fmt.Println("Lista despues: ", p.Likes)

		} else {
			p.Likes = likes // remains the same

		}

		// We don't change the rest of the attributes
		p.Path = path
		p.Comments = comments
		p.Date = date
		p.UserId = userId

		res, err := db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
			p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
		if err != nil {
			return p, err
		}

		fmt.Println(res)

	}

	return p, nil

}
