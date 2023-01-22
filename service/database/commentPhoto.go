package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

func (db *appdbimpl) CommentPhoto(c Comment, p Photo, u User) (Comment, Photo, User, error) {

	var photoId, profilePic, photoUserId int
	var userName, userNameTarget, followers, banned, photos, likes, photosComments, path string
	var date time.Time
	// var userId int

	//search for the user that comments
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, c.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if userName == "" {
		return c, p, u, errors.New("User not found")
	}

	//then we search the photo id. If it doesn't exist, we cannot comment on the photo

	rows2, err := db.c.Query(`select id,userid,path,likes,comments,date from photos where id=?`, c.PhotoId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&photoId, &photoUserId, &path, &likes, &photosComments, &date)

		if err != nil {
			log.Fatal(err)
		}

	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	//lastly, we need to check up the user that gets commented
	p.UserId = photoUserId //to check for the target id

	fmt.Println("P. userid: ", p.UserId)

	rows3, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows3.Close()

	for rows3.Next() {

		err := rows3.Scan(&userNameTarget, &profilePic, &followers, &banned, &photos)

		if err != nil {
			log.Fatal(err)
		}

	}

	err = rows3.Err()
	if err != nil {
		log.Fatal(err)
	}

	if userNameTarget == "" {
		return c, p, u, errors.New("Target user not found")
	}

	if userName != "" && userNameTarget != "" && photoId != 0 { //comment has not been uploaded before and the user and photo exist
		c.Date = time.Now()

		res, err := db.c.Exec(`INSERT INTO comments (commentid,content,photoid,userid,date) VALUES (NULL,?,?,?,?)`,
			c.ID, c.Content, c.PhotoId, c.UserId, c.Date)
		if err != nil {
			return c, p, u, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return c, p, u, err
		}

		c.ID = int(lastInsertID)

		// We also have to update the comments's stream of the COMMENTED user
		u.Name = userNameTarget
		u.ID = p.UserId //this is important: the id of the user we need to update is not the one who comments, but the one who gets the comment on the photo
		u.Followers = followers
		u.Banned = banned
		u.ProfilePic = profilePic

		//and the photos' comment
		p.UserId = photoUserId
		p.Likes = likes
		p.Date = date
		p.Path = path

		// UPDATING the photo

		var add string

		new_list := photosComments[0 : len(photosComments)-1]
		newComment := `{"User":"` + u.Name + `", "Comment":"` + c.Content + `"}`
		if photosComments == "[]" {
			add = newComment + "]"

		} else {
			add = "," + newComment + "]"

		}
		new_list += add

		p.Comments = new_list

		// UPDATING the photo's

		// We cast the photos to a Photo's array, then we change the one who gets commented
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
				castPhotos[i].Comments = p.Comments
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

		res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
			u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
		if err != nil {
			return c, p, u, err
		}

		return c, p, u, nil
	}
	return c, p, u, errors.New("Comment already uploaded")

}
