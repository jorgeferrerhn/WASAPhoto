package database

import (
	"fmt"
	"log"
	"strings"
)

func (db *appdbimpl) LikePhoto(p Photo) (Photo, error) {

	var likes string
	var userName string

	//search for the user
	rows, err := db.c.Query(`select name from users where id=?`, p.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userName)

		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//search the photo
	rows2, err := db.c.Query(`select likes from photos where id=?`, p.ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err := rows2.Scan(&likes)

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(likes)
	}

	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	if userName != "" { //user exists

		lista := make([]int, 0)

		fmt.Println("likes: ", likes)

		if likes != "[]" {
			fmt.Println("Likes: ", likes)
			output := likes[1 : len(likes)-1]
			res := strings.Split(output, ",")
			fmt.Println("This output: ", res)
		}

		lista = append(lista, p.UserId)

		// cast from array to string
		justString := fmt.Sprint(lista)

		//pass to database
		p.Likes = justString

		fmt.Println("Lista despues: ", p.Likes)

		res, err := db.c.Exec(`UPDATE photos SET likes=? WHERE id=?`,
			p.Likes, p.ID)
		if err != nil {
			return p, err
		}

		fmt.Println(res)

		/*

			c.Date = time.Now()




			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return c, err
			}

			c.ID = uint64(lastInsertID)

		*/

	}

	return p, nil

}
