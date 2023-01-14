package database

import (
	"fmt"
	"log"
	"strings"
)

func (db *appdbimpl) UnbanUser(id1 int, id2 int) (int, error) {

	var followers, banned string
	//search the photo
	rows, err := db.c.Query(`select banned from users where id=?`, id2)

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

	fmt.Println("Banned: ", banned)

	if followers != "[]" {
		output := banned[1 : len(banned)-1]
		res := strings.Split(output, ",")
		fmt.Println(res)
	}

	lista = append(lista, id1)
	fmt.Println(lista)

	//actualizar base de datos de usuarios (delete)

	//también tenemos que comprobar que el usuario ya estaba baneado

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//update list of followers

	return 0, nil

}
