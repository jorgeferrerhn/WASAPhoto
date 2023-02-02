package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (db *appdbimpl) UploadPhoto(p Photo, u User) (Photo, User, error) {

	var photoId, profilePic int
	var userName, followers, banned, photos string

	// search for the user
	rows2, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {
		return p, u, err
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&userName, &profilePic, &followers, &banned, &photos)

		if err != nil {
			return p, u, err
		}
	}

	err = rows2.Err()
	if err != nil {
		return p, u, err
	}

	if userName == "" {
		return p, u, errors.New("User not found")
	}

	// search for the photo id (check if it existed)
	rows, err := db.c.Query(`select id from photos where path=? and userid=?`, p.Path, p.UserId)

	if err != nil {
		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&photoId)

		if err != nil {
			return p, u, err
		}
	}

	err = rows.Err()
	if err != nil {
		return p, u, err
	}

	if photoId != 0 {
		return p, u, errors.New("Photo already uploaded")
	}

	if p.Path == "" {
		return p, u, errors.New("Request body Empty")
	}

	p.Likes = "[]"
	p.Comments = "[]"
	p.Date = time.Now()

	// We upload the photo and insert it to the photo's database
	res, err := db.c.Exec(`INSERT INTO photos (id,userid,path,likes,comments,date) VALUES (NULL,?,?,?,?,?)`,
		p.UserId, p.Path, p.Likes, p.Comments, p.Date)
	if err != nil {
		return p, u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return p, u, err
	}

	p.ID = int(lastInsertID)

	//  We also have to update the photo's stream of the user
	u.Name = userName
	u.Followers = followers
	u.Banned = banned
	u.ProfilePic = profilePic

	// Here, we have to take the photos and cast them to {1, 1, ... } --> json.Unmarshal
	in2 := []byte(photos)
	var castPhotos []Photo
	err = json.Unmarshal(in2, &castPhotos)
	if err != nil {
		return p, u, err
	}

	castPhotos = append(castPhotos, p)

	fmt.Println(castPhotos)

	// Here, we have to store the photo as {"ID": 1, "UserID": ...} -->json.Marshal
	savePhotos, err := json.Marshal(castPhotos)
	if err != nil {
		return p, u, err
	}

	fmt.Println("Saved photos: ", string(savePhotos))
	u.Photos = string(savePhotos)

	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return p, u, err
	}

	return p, u, nil

}
