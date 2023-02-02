package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) UnbanUser(user1 User, user2 User) (User, error) {

	//  search for the user that follows
	rows, err := db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user1.ID)

	if err != nil {
		return user1, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&user1.Name, &user1.ProfilePic, &user1.Followers, &user1.Banned, &user1.Photos)

		if err != nil {
			return user1, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user1, err
	}

	if user1.Name == "" {
		return user1, errors.New("User 1 not found")
	}

	//  search for the user that get followed
	rows, err = db.c.Query(`SELECT name FROM users WHERE id=?`, user2.ID)

	if err != nil {
		return user1, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&user2.Name)

		if err != nil {
			return user1, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user1, err
	}

	if user2.Name == "" {
		return user1, errors.New("User 2 not found")
	}

	banned := strings.Contains(user1.Banned, fmt.Sprint(user2.ID))
	if !banned {
		return user1, errors.New("User2 wasn't previously banned by user1")
	}
	// We cast to list the string
	newList := "["
	counter := 1
	//  Updating the followers' list
	for i := 1; i < len(user1.Banned)-1; i++ { // Chapuza: esto hay que cambiarlo
		c := string(user1.Banned[i]) //  rune to string
		if c == "," {
			number := user1.Banned[counter:i] //  takes up to that position
			if number != fmt.Sprint(user1.ID) {
				newList += number + ","
			}
		}
	}

	newList = newList[:len(newList)-1] + "]" //  extract the last comma and add the last bracket
	if newList == "]" {
		//  It was empty
		newList = "[]"
	}

	user1.Banned = newList

	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user1.Name, user1.ProfilePic, user1.Followers, user1.Banned, user1.Photos, user1.ID)
	if err != nil {
		return user1, errors.New("Error in " + fmt.Sprint(res))
	}

	return user1, nil

}
