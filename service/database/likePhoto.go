package database

import (
	"fmt"
	"log"
	"strings"
)

func (db *appdbimpl) LikePhoto(photoId int) (int, error) {

	var uID int
	var likes string

	//search the photo
	rows, err := db.c.Query(`select userid,likes from photos where id=?`, photoId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&uID, &likes)

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(likes)
	}

	lista := make([]int, 0)

	fmt.Println("likes: ", likes)

	if likes != "[]" {
		output := likes[1 : len(likes)-1]
		res := strings.Split(output, ",")
		fmt.Println(res)
	}

	lista = append(lista, uID)
	fmt.Println(lista)

	//actualizar base de datos de fotos

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//update number of likes

	return 0, nil

}
