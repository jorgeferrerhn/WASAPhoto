package database

import (
	"errors"
	"fmt"
)

func (db *appdbimpl) SetMyUserName(u User) (User, error) {

	// we get all the information from the user
	var profilePic, searchId int
	var userNameTarget, followers, banned, photos string

	//first we search the user. It should have a unique username, so we'll search for it
	rows, err := db.c.Query(`select name,profilepic,followers,banned,photos from users where id=?`, u.ID)

	if err != nil {
		return u, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userNameTarget, &profilePic, &followers, &banned, &photos)

		if err != nil {
			return u, err
		}

	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if userNameTarget == "" { //invalid user id
		return u, errors.New("This user doesn't exist!")
	}
	// Then, we have to check if the username has been already taken by everyone else
	rows2, err := db.c.Query(`select id from users where name=?`, u.Name)

	for rows2.Next() {

		err := rows2.Scan(&searchId)

		if err != nil {
			return u, err
		}

	}
	err = rows.Err()
	if err != nil {
		return u, err
	}

	if searchId != 0 { // There's someone with that username already
		return u, errors.New("This username is already picked!")
	}

	// update the information
	u.ProfilePic = profilePic
	u.Followers = followers
	u.Banned = banned
	u.Photos = photos

	res, err := db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		u.Name, u.ProfilePic, u.Followers, u.Banned, u.Photos, u.ID)

	if err != nil {
		return u, errors.New("Error in " + fmt.Sprint(res))
	}

	return u, nil
}
