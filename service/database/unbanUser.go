package database

func (db *appdbimpl) UnbanUser(id1 int, id2 int) (int, error) {

	var followers, banned string
	//search the photo
	rows, err := db.c.Query(`select banned,followers from users where id=?`, id2)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&banned, &followers)

		if err != nil {

			return 0, err
		}

	}

	lista := make([]int, 0)

	lista = append(lista, id1)

	//actualizar base de datos de usuarios (delete)

	//también tenemos que comprobar que el usuario ya estaba baneado

	err = rows.Err()
	if err != nil {
		return 0, err
	}

	//update list of followers

	return 0, nil

}
