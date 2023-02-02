package database

import (
	"errors"
	"time"
)

func (db *appdbimpl) UploadLogo(p Photo, u User) (Photo, User, error) {

	var photoId int

	// search for the user
	rows2, err2 := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, p.UserId)

	if err2 != nil {
		return p, u, err2
	}

	defer rows2.Close()

	for rows2.Next() {

		err2 = rows2.Scan(&u.Name, &u.ProfilePic, &u.Followers, &u.Banned, &u.Photos)

		if err2 != nil {
			return p, u, err2
		}
	}

	err2 = rows2.Err()
	if err2 != nil {
		return p, u, err2
	}

	if u.Name == "" {
		return p, u, errors.New("User not found")
	}

	// search for the photo id (check if it existed)
	rows, err := db.c.Query(`select id from photos where path=? and userid=?`, p.Path, p.UserId)

	if err != nil {
		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&photoId)

		if err != nil {
			return p, u, err
		}
	}

	err = rows.Err()
	if err != nil {
		return p, u, err
	}

	if photoId == u.ProfilePic && photoId != 0 {
		return p, u, errors.New("This is the current profile picture!")
	}

	if p.Path == "" {
		return p, u, errors.New("Request body Empty")
	}

	p.Likes = "[]"
	p.Comments = "[]"
	p.Date = time.Now()

	//  We upload the photo and insert it to the photo's database
	res, e := db.c.Exec(`INSERT INTO photos (id,userid,path,likes,comments,date) VALUES (NULL,?,?,?,?,?)`,
		p.UserId, p.Path, p.Likes, p.Comments, p.Date)
	if e != nil {
		return p, u, e
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return p, u, err
	}

	p.ID = int(lastInsertID)

	//  We also have to update the profile picture ID
	u.ProfilePic = p.ID

	if err != nil {
		return p, u, err
	}

	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)
	if err != nil {
		return p, u, err
	}

	return p, u, nil

}
