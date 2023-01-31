package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (db *appdbimpl) UncommentPhoto(c Comment, p Photo, u User) (Comment, Photo, User, error) {

	var profilePic, photoUserId, pcId, ucId int
	var userName, userNameTarget, followers, banned, photos, likes, photosComments, path string
	var content, dateComment string
	var date time.Time
	// var userId int

	//search for the user that comments
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, c.UserId)

	if err != nil {
		return c, p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName)

		if err != nil {
			return c, p, u, err
		}
	}

	err = rows.Err()
	if err != nil {
		return c, p, u, err
	}

	if userName == "" {
		return c, p, u, errors.New("User not found")
	}

	//then we search the comment ID. If it doesn't exist, we cannot uncomment the photo

	rows2, err := db.c.Query(`select content,photoid,userid,date from comments where commentid=?`, c.ID)

	if err != nil {
		return c, p, u, err
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&content, &pcId, &ucId, &dateComment)

		if err != nil {
			return c, p, u, err
		}

	}

	err = rows2.Err()
	if err != nil {
		return c, p, u, err
	}

	if content == "" || pcId == 0 || ucId == 0 {
		return c, p, u, errors.New("Comment not found")
	}

	p.ID = pcId // We update the photoId

	//lastly, we need to check if the photo previously existed
	rows3, err := db.c.Query(`select userid,path,likes,comments,date from photos where id=?`, p.ID)

	if err != nil {
		return c, p, u, err
	}

	defer rows3.Close()

	for rows3.Next() {

		err := rows3.Scan(&photoUserId, &path, &likes, &photosComments, &date)

		if err != nil {
			return c, p, u, err
		}

	}

	err = rows3.Err()
	if err != nil {
		return c, p, u, err
	}

	if path == "" {
		return c, p, u, errors.New("Photo didn't exist")
	}

	p.UserId = photoUserId
	p.Path = path
	p.Likes = likes
	// p.Comments = photosComments
	p.Date = date

	// Delete comment from comments' table

	res, err := db.c.Exec(`DELETE FROM comments WHERE commentid=?`, c.ID)
	if err != nil {
		return c, p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

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
	/*
		var add string

		for i := 1; i < len(photosComments)-1; i++ {

			c := string(likes[i]) // rune to string

			if c == "," {
				number := likes[counter:i] // takes up to that position
				if number != fmt.Sprint(u.ID) {
					newList += number + ","
				}
			}
		}

		new_list := photosComments[0 : len(photosComments)-1]
		newComment := `['User':'` + fmt.Sprint(u.ID) + `', 'Comment':'` + c.Content + `']`
		if photosComments == "[]" {
			add = newComment + "]"

		} else {
			add = "," + newComment + "]"

		}
		new_list += add

		p.Comments = new_list

	*/

	// UPDATING the photo's

	// We cast the photos to a Photo's array, then we change the one who gets commented
	in := []byte(photos)
	var castPhotos []Photo
	err = json.Unmarshal(in, &castPhotos)
	if err != nil {

	}

	newPhotos := "["
	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { //this is the one who gets commented

			castPhotos[i].Comments = p.Comments
		}
		newPhoto := `{"id": ` + fmt.Sprint(castPhotos[i].ID) + `, "userid": ` + fmt.Sprint(castPhotos[i].UserId) + `, "path": "` + castPhotos[i].Path + `", "likes": "` + castPhotos[i].Likes + `", "comments": "` + castPhotos[i].Comments + `", "date": "` + castPhotos[i].Date.Format(time.RFC3339) + `"}`
		if i == len(castPhotos)-1 {
			newPhotos += newPhoto + "]"
		} else {
			newPhotos += newPhoto + ","
		}
	}

	u.Photos = newPhotos

	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return c, p, u, err
	}

	return c, p, u, nil

}
