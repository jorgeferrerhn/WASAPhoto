package database

import (
	"database/sql"

	"errors"
	"fmt"
	"log"
	"strings"
)

func (db *appdbimpl) FollowUser(user1 User, user2 User) (User, error) {
	var name2, followers2, banned2, photos2, name1 string
	var profilePic2 int

	// We have to check if both users exist

	// search for the user that follows
	rows, err := db.c.Query(`SELECT name FROM users WHERE id=?`, user1.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name1)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if name1 == "" {
		return user2, errors.New("Follower not found")
	}

	// search for the user that get followed
	rows, err = db.c.Query(`SELECT name,profilepic,followers,banned,photos FROM users WHERE id=?`, user2.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&name2, &profilePic2, &followers2, &banned2, &photos2)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if name2 == "" {
		return user2, errors.New("Followed not found")
	}

	followed := strings.ContainsAny(followers2, fmt.Sprint(user1.ID))

	if !followed {
		var add string

		new_list := followers2[0 : len(followers2)-1]

		if followers2 == "[]" {
			add = fmt.Sprint(user1.ID) + "]"

		} else {
			add = "," + fmt.Sprint(user1.ID) + "]"

		}
		new_list += add

		user2.Followers = new_list

	} else {
		user2.Followers = followers2 // remains the same

	}

	fmt.Println("Followers of user2 after: ", user2.Followers)

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
