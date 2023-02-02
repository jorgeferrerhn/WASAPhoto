package database

import (
	"errors"
)

func (db *appdbimpl) GetLogo(p Photo, u User) (Photo, User, error) {

	rows, err := db.c.Query(`select id, name, followers,banned,photos from users where profilepic=?`, u.ProfilePic)

	if err != nil {
		return p, u, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Followers, &u.Banned, &u.Photos)

		if err != nil {
			return p, u, err
		}

	}

	if u.Name == "" {
		return p, u, errors.New("User didn't exist!")
	}
	err = rows.Err()
	if err != nil {
		return p, u, err
	}

	defer rows.Close()

	rows2, err2 := db.c.Query(`select userid,path,likes,comments,date from photos where id=?`, p.ID)

	if err2 != nil {
		return p, u, err2
	}
	for rows2.Next() {
		err = rows2.Scan(&p.UserId, &p.Path, &p.Likes, &p.Comments, &p.Date)

		if err != nil {
			return p, u, err
		}

	}
	err = rows2.Err()
	if err != nil {
		return p, u, err
	}

	return p, u, nil
}
