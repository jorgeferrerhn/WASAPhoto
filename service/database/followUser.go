package database

import (
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

		err2 := rows.Scan(&user1.Name)

		if err2 != nil {
			return user2, err2
		}
	}

	err3 := rows.Err()
	if err3 != nil {
		return user2, err3
	}

	if user1.Name == "" {
		return user2, errors.New("Follower not found")
	}

	// search for the user that get followed
	rows2, err4 := db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user2.ID)

	if err4 != nil {
		return user2, err4
	}

	defer rows2.Close()

	for rows2.Next() {

		err5 := rows2.Scan(&user2.Name, &user2.ProfilePic, &user2.Followers, &user2.Banned, &user2.Photos)

		if err5 != nil {
			return user2, err5
		}
	}

	err6 := rows2.Err()
	if err6 != nil {
		return user2, err6
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

	res2, err7 := db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user2.Name, user2.ProfilePic, user2.Followers, user2.Banned, user2.Photos, user2.ID)
	if err7 != nil {
		return user2, errors.New("Error in " + fmt.Sprint(res2))
	}

	return user2, nil

}
