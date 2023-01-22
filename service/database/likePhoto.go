package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func (db *appdbimpl) LikePhoto(p Photo, u User) (Photo, User, error) {

	var likes, path, comments, followers, banned, photos string
	var userId, profilePic int
	var date time.Time
	var userName string

	//search for the user
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {

		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName, &profilePic, &followers, &banned, &photos)

		if err != nil {
			return p, u, err
		}
	}

	if userName == "" {
		//el usuario no existía
		return p, u, errors.New("User not found")
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
		return p, u, errors.New("Photo not found")
	}

	var liked bool

	// We check that the photo hasn't been liked before
	strId := fmt.Sprint(p.UserId)
	liked = strings.ContainsAny(likes, strId)
	fmt.Println("Likes : ", likes, " of ", p.UserId, ": ", liked)

	fmt.Println("Liked: ", liked)

	if !liked {
		var add string

		new_list := likes[0 : len(likes)-1]

		if likes == "[]" {
			add = userName + "]"

		} else {
			add = "," + userName + "]"

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

	//We update the user's photos and the photos' stream

	// We cast the photos to a Photo's array, then we change the one who gets commented
	fmt.Println("Photos: ", photos)
	in := []byte(photos)
	var castPhotos []Photo
	err = json.Unmarshal(in, &castPhotos)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Cast photos: ", castPhotos)
	newPhotos := "["
	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { //this is the one who gets commented
			fmt.Println("HERE")
			castPhotos[i].Likes = p.Likes
		}
		newPhoto := `{"id": ` + fmt.Sprint(castPhotos[i].ID) + `, "userid": ` + fmt.Sprint(castPhotos[i].UserId) + `, "path": "` + castPhotos[i].Path + `", "likes": "` + castPhotos[i].Likes + `", "comments": "` + castPhotos[i].Comments + `", "date": "` + castPhotos[i].Date.Format(time.RFC3339) + `"}`
		if i == len(castPhotos)-1 {
			newPhotos += newPhoto + "]"
		} else {
			newPhotos += newPhoto + ","
		}
	}
	fmt.Println("New photos: ", newPhotos)
	u.Photos = newPhotos

	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return p, u, err
	}

	res, err = db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
		p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
	if err != nil {
		return p, u, err
	}

	fmt.Println(res)

	return p, u, nil

}
