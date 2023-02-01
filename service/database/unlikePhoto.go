package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func (db *appdbimpl) UnlikePhoto(p Photo, u User) (Photo, User, error) {

	var likes, path, comments, followers, banned, photos string
	var userId, profilePic int
	var date time.Time
	var userName string

	// search for the user
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {

		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName, &profilePic, &followers, &banned, &photos)

		if err != nil {
			return p, u, err
		}
	}

	if userName == "" {
		// el usuario no existía
		return p, u, errors.New("User not found")
	}
	err = rows.Err()
	if err != nil {
		return p, u, err
	}

	// search the photo
	rows2, err := db.c.Query(`select userId,path,likes,comments,date from photos where id=?`, p.ID)

	if err != nil {
		return p, u, err
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&userId, &path, &likes, &comments, &date)

		if err != nil {

			return p, u, err
		}

	}

	err = rows2.Err()
	if err != nil {
		return p, u, err
	}

	if path == "" {
		// el usuario no existía
		return p, u, errors.New("Photo not found")
	}

	liked := strings.Contains(likes, fmt.Sprint(u.ID))

	if !liked {
		return p, u, errors.New("Photo wasn't previously liked!")
	}

	// We cast to list the string

	newList := "["
	counter := 1

	//  Updating the followers' list
	for i := 1; i < len(likes)-1; i++ {

		c := string(likes[i]) //  rune to string

		if c == "," {
			number := likes[counter:i] //  takes up to that position
			if number != fmt.Sprint(u.ID) {
				newList += number + ","
			}
		}
	}

	newList = newList[:len(newList)-1] + "]" //  extract the last comma and add the last bracket
	if newList == "]" {
		//  It was empty
		newList = "[]"
	}

	p.Likes = newList // update the likes list
	// We don't change the rest of the attributes
	p.Path = path
	p.Comments = comments
	p.Date = date
	p.UserId = userId

	//We update the user's photos and the photos' stream

	// We cast the photos to a Photo's array, then we change the one who gets commented

	in := []byte(photos)
	var castPhotos []Photo
	err = json.Unmarshal(in, &castPhotos)
	if err != nil {

	}

	newPhotos := "["
	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { //this is the one who gets liked

			castPhotos[i].Likes = p.Likes
		}
		newPhoto := `{"id": ` + fmt.Sprint(castPhotos[i].ID) + `,"userid": ` + fmt.Sprint(castPhotos[i].UserId) + `,"path": "` + castPhotos[i].Path + `","likes": "` + castPhotos[i].Likes + `","comments": "` + castPhotos[i].Comments + `","date": "` + castPhotos[i].Date.Format(time.RFC3339) + `"}`
		//newPhoto := fmt.Sprint(p)

		if i == len(castPhotos)-1 {
			newPhotos += newPhoto + "]"
		} else {
			newPhotos += newPhoto + ","
		}
	}

	u.Name = userName
	u.ID = p.UserId //this is important: the id of the user we need to update is not the one who likes, but the one who gets the like on the photo
	u.Followers = followers
	u.Banned = banned
	u.ProfilePic = profilePic
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
		return p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	return p, u, nil

}
