package database

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func (db *appdbimpl) CommentPhoto(c Comment, p Photo, u User) (Comment, Photo, User, error) {

	var photoId, profilePic, photoUserId int
	var userName, followers, banned, photos, likes, photosComments, path string
	var date time.Time
	// var userId int

	//search for the user
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, c.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName, &profilePic, &followers, &banned, &photos)

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

	if path == "" {
		return c, p, u, errors.New("Photo not found")
	}

	if userName != "" && photoId != 0 { //comment has not been uploaded before and the user and photo exist
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

		// We also have to update the comments's stream of the user
		u.Followers = followers
		u.Photos = photos
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
		newComment := "{'User':" + u.Name + " Comment:" + c.Content + "}"
		if photosComments == "[]" {
			add = newComment + "]"

		} else {
			add = "," + newComment + "]"

		}
		new_list += add

		p.Comments = new_list

		fmt.Println("Lista despues: ", p.Comments)

		res, err = db.c.Exec(`UPDATE photos SET userid=?,path=?,likes=?,comments=?,date=? WHERE id=?`,
			p.UserId, p.Path, p.Likes, p.Comments, p.Date, p.ID)
		if err != nil {
			return c, p, u, err
		}

		fmt.Println(res)

		// UPDATING the user's stream
		//res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		//	u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)

		return c, p, u, nil
	}
	return c, p, u, errors.New("Comment already uploaded")

}
