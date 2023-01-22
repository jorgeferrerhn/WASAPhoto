package database

func (db *appdbimpl) UnlikePhoto(photoId int) (int, error) {

	var uID int
	var likes string

	//search the photo
	rows, err := db.c.Query(`select userid,likes from photos where id=?`, photoId)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&uID, &likes)

		if err != nil {

			return 0, err
		}

	}

	lista := make([]int, 0)

	lista = append(lista, uID)

	//actualizar base de datos de fotos

	err = rows.Err()
	if err != nil {
		return 0, err
	}

	//update number of likes

	return 0, nil

}
