package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (db *appdbimpl) UnfollowUser(user1 User, user2 User) (User, error) {
	var name2, followers2, banned2, photos2, name1 string
	var profilePic2 int

	// We have to check if both users exist

	// search for the user that follows
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, user1.ID)

	if err != nil {
		return user2, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name1)

		if err != nil {
			return user2, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user2, err
	}

	if name1 == "" {
		return user2, errors.New("Follower not found")
	}

	// search for the user that get followed
	rows, err = db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user2.ID)

	if err != nil {
		return user2, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name2, &profilePic2, &followers2, &banned2, &photos2)

		if err != nil {
			return user2, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user2, err
	}

	if name2 == "" {
		return user2, errors.New("Followed not found")
	}

	followed := strings.Contains(followers2, fmt.Sprint(user1.ID))

	if !followed {
		return user2, errors.New("User2 wasn't previously followed by user1")
	}

	//We cast to list the string

	newList := "["
	counter := 1

	// Updating the followers' list
	for i := 1; i < len(followers2)-1; i++ {

		c := string(followers2[i]) // rune to string

		if c == "," {
			number := followers2[counter:i] // takes up to that position
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
	user2.Followers = newList
	user2.Name = name2
	user2.Banned = banned2
	user2.ProfilePic = profilePic2
	user2.Photos = photos2

	var res sql.Result
	res, err = db.c.Exec(`UPDATE users SET name=?,profilepic=?,followers=?,banned=?,photos=? WHERE id=?`,
		user2.Name, user2.ProfilePic, user2.Followers, user2.Banned, user2.Photos, user2.ID)
	if err != nil {
		return user2, errors.New("Error in " + fmt.Sprint(res))
	}

	//update list of followers

	return user2, nil

}
