package database

import (
	"fmt"
	"log"
	"strings"
)

func (db *appdbimpl) UnfollowUser(id1 int, id2 int) (int, error) {

	var followers string
	//search the photo
	rows, err := db.c.Query(`select followers from users where id=?`, id2)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&followers)

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(followers)
	}

	lista := make([]int, 0)

	fmt.Println("Followers: ", followers)

	if followers != "[]" {
		output := followers[1 : len(followers)-1]
		res := strings.Split(output, ",")
		fmt.Println(res)
	}

	lista = append(lista, id1)
	fmt.Println(lista)

	//actualizar base de datos de usuarios (delete)

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//update list of followers

	return 0, nil

}
