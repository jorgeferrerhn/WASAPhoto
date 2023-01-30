package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) BanUser(user1 User, user2 User) (User, error) {

	var name1, followers1, banned1, photos1, name2 string
	var profilePic1 int

	// We have to check if both users exist

	// search for the user that follows
	rows, err := db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user1.ID)

	if err != nil {
		return user1, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name1, &profilePic1, &followers1, &banned1, &photos1)

		if err != nil {
			return user1, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user1, err
	}

	if name1 == "" {
		return user1, errors.New("User 1 not found")
	}

	// search for the user that get followed
	rows, err = db.c.Query(`SELECT name FROM users WHERE id=?`, user2.ID)

	if err != nil {
		return user1, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name2)

		if err != nil {
			return user1, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user1, err
	}

	if name2 == "" {
		return user1, errors.New("User 2 not found")
	}

	banned := strings.ContainsAny(banned1, fmt.Sprint(user2.ID))

	if !banned {
		var add string

		newList := banned1[0 : len(banned1)-1]

		if banned1 == "[]" {
			add = fmt.Sprint(user2.ID) + "]"

		} else {
			add = "," + fmt.Sprint(user2.ID) + "]"

		}
		newList += add

		user1.Banned = newList

	} else {
		user1.Banned = banned1 // remains the same

	}

	user1.Name = name1
	user1.ProfilePic = profilePic1
	user1.Photos = photos1
	user1.Followers = followers1
	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user1.Name, user1.ProfilePic, user1.Followers, user1.Banned, user1.Photos, user1.ID)
	if err != nil {
		return user1, errors.New("Error in " + fmt.Sprint(res))
	}

	//update list of followers

	return user1, nil

}
