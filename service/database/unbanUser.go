package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) UnbanUser(user1 User, user2 User) (User, error) {
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

	banned := strings.Contains(banned1, fmt.Sprint(user2.ID))

	if !banned {
		return user1, errors.New("User2 wasn't previously banned by user1")
	}

	//We cast to list the string

	newList := "["
	counter := 1

	// Updating the followers' list
	for i := 1; i < len(banned1)-1; i++ {

		c := string(banned1[i]) // rune to string

		if c == "," {
			number := banned1[counter:i] // takes up to that position
			if number != fmt.Sprint(user1.ID) {
				newList += number + ","
			}
		}
	}

	newList = newList[:len(newList)-1] + "]" // extract the last comma and add the last bracket
	if newList == "]" {
		// It was empty
		newList = "[]"
	}
	user1.Banned = newList
	user1.Name = name1
	user1.Followers = followers1
	user1.ProfilePic = profilePic1
	user1.Photos = photos1

	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user1.Name, user1.ProfilePic, user1.Followers, user1.Banned, user1.Photos, user1.ID)
	if err != nil {
		return user1, errors.New("Error in " + fmt.Sprint(res))
	}

	//update list of followers

	return user1, nil

}
