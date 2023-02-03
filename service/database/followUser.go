package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) FollowUser(user1 User, user2 User) (User, error) {
	// search for the user that follows
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, user1.ID)

	if err != nil {
		return user2, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&user1.Name)

		if err != nil {
			return user2, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user2, err
	}

	if user1.Name == "" {
		return user2, errors.New("Follower not found")
	}

	// search for the user that get followed
	rows, err = db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user2.ID)

	if err != nil {
		return user2, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&user2.Name, &user2.ProfilePic, &user2.Followers, &user2.Banned, &user2.Photos)

		if err != nil {
			return user2, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user2, err
	}

	if user2.Name == "" {
		return user2, errors.New("Followed not found")
	}

	followed := strings.Contains(user2.Followers, fmt.Sprint(user1.ID))

	if !followed { // Chapuza: esto hay que cambiarlo
		var add string
		newList := user2.Followers[0 : len(user2.Followers)-1]
		if user2.Followers == "[]" {
			add = fmt.Sprint(user1.ID) + "]"
		} else {
			add = "," + fmt.Sprint(user1.ID) + "]"
		}
		newList += add
		user2.Followers = newList

	}

	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user2.Name, user2.ProfilePic, user2.Followers, user2.Banned, user2.Photos, user2.ID)
	if err != nil {
		return user2, errors.New("Error in " + fmt.Sprint(res))
	}

	return user2, nil

}
