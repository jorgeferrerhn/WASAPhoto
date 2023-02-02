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
	//  var userId int

	// search for the user that comments
	rows, err := db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, c.UserId)

	if err != nil {
		return c, p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName, &profilePic, &followers, &banned, &photos)

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

	// then we search the comment ID. If it doesn't exist, we cannot uncomment the photo

	// fmt.println("Hasta aqui llega")

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

	p.ID = pcId //  We update the photoId

	// lastly, we need to check if the photo previously existed
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

	// Here we cast the comments to "raw format" { 1 Content ...} --> json.Unmarshal

	in := []byte(photosComments)
	var castComments []Comment
	err = json.Unmarshal(in, &castComments)
	if err != nil {
		return c, p, u, err
	}

	// We create a new comments array, so we only save the wanted comments
	var newComments []Comment

	for i := 0; i < len(castComments); i++ {

		if castComments[i].ID != c.ID { // we add everything except the comments
			newComments = append(newComments, castComments[i])
		}

	}

	// Now, in newComments we have only the comments we want. We have to store them as {"ID": 1, "Content": ...} --> json.Marshal
	saveComments, err := json.Marshal(newComments)
	p.Comments = string(saveComments)

	res, err = db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
		p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
	if err != nil {
		return c, p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	// Here we update the information of the photo on "raw format" { 1 Content ...} --> json.Unmarshal
	in2 := []byte(photos)
	var castPhotos []Photo
	err = json.Unmarshal(in2, &castPhotos)
	if err != nil {
		return c, p, u, err
	}

	for i := 0; i < len(castPhotos); i++ {
		if castPhotos[i].ID == p.ID { //this is the one who gets commented
			castPhotos[i].Comments = p.Comments
		}
	}
	savePhotos, err := json.Marshal(castPhotos)
	u.Photos = string(savePhotos)

	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return c, p, u, err
	}

	return c, p, u, nil

}
