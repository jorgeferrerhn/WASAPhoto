package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (db *appdbimpl) CommentPhoto(c Comment, p Photo, u User) (Comment, Photo, User, error) {

	var userNameTarget string
	var castComments []Comment
	var castPhotos []Photo

	// search for the user that comments
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, c.UserId)

	if err != nil {
		return c, p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&u.Name)

		if err != nil {
			return c, p, u, err
		}
	}

	err = rows.Err()
	if err != nil {
		return c, p, u, err
	}

	if u.Name == "" {
		return c, p, u, errors.New("User not found")
	}

	// then we search the photo id. If it doesn't exist, we cannot comment on the photo

	rows2, err2 := db.c.Query(`select id,userid,path,likes,comments,date from photos where id=?`, c.PhotoId)

	if err2 != nil {
		return c, p, u, err2
	}

	defer rows2.Close()

	for rows2.Next() {

		err = rows2.Scan(&p.ID, &p.UserId, &p.Path, &p.Likes, &p.Comments, &p.Date)

		if err != nil {
			return c, p, u, err
		}

	}

	err = rows2.Err()
	if err != nil {
		return c, p, u, err
	}
	if p.ID == 0 {
		return c, p, u, errors.New("Photo not found")
	}

	// lastly, we need to check up the user that gets commented

	rows3, err3 := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err3 != nil {
		return c, p, u, err3
	}

	defer rows3.Close()

	for rows3.Next() {

		err3 = rows3.Scan(&userNameTarget, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err3 != nil {
			return c, p, u, err
		}

	}

	err3 = rows3.Err()
	if err3 != nil {
		return c, p, u, err
	}

	if userNameTarget == "" {
		return c, p, u, errors.New("Target user not found")
	}

	c.Date = time.Now()

	res, e := db.c.Exec(`INSERT INTO comments (commentid,content,photoid,userid,date) VALUES (NULL,?,?,?,?)`, c.Content, c.PhotoId, c.UserId, c.Date)
	if e != nil {
		return c, p, u, e
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return c, p, u, err
	}

	c.ID = int(lastInsertID)

	//  We also have to update the comments's stream of the COMMENTED user
	u.Name = userNameTarget
	u.ID = p.UserId // this is important: the id of the user we need to update is not the one who comments, but the one who gets the comment on the photo

	//  UPDATING the photo

	// Here we append the comment on "raw format" { 1 Content ...} --> json.Unmarshal
	in := []byte(p.Comments)
	err = json.Unmarshal(in, &castComments)
	if err != nil {
		return c, p, u, err
	}

	castComments = append(castComments, c)

	// Here we save the comment photo as {"ID": 1,"Content": ...} --> json.Marshal
	saveComments, err := json.Marshal(castComments)
	p.Comments = string(saveComments)

	// UPDATING the photo's

	res, err = db.c.Exec(`UPDATE photos SET path=?,comments=?,date=?,userid=?,likes=? WHERE id=?`,
		p.Path, p.Comments, p.Date, p.UserId, p.Likes, p.ID)
	if err != nil {
		return c, p, u, errors.New("Error in: " + fmt.Sprint(res))
	}

	// Here we update the information of the photo on "raw format" { 1 Content ...} --> json.Unmarshal
	in2 := []byte(u.Photos)
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
	if err != nil {
		return c, p, u, err
	}
	u.Photos = string(savePhotos)

	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return c, p, u, err
	}

	return c, p, u, nil

}
