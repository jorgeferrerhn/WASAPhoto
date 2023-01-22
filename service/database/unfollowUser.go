package database

func (db *appdbimpl) UnfollowUser(id1 int, id2 int) (int, error) {

	var followers string
	//search the photo
	rows, err := db.c.Query(`select followers from users where id=?`, id2)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&followers)

		if err != nil {

			return 0, err
		}

	}

	lista := make([]int, 0)

	lista = append(lista, id1)

	//actualizar base de datos de usuarios (delete)

	err = rows.Err()
	if err != nil {
		return 0, err
	}

	//update list of followers

	return 0, nil

}
