/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.
To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.
For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	//  Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

//		User struct represent a fountain in every API call between this package and the outside world.
//
//	 Note that the internal representation of fountain in the database might be different.
type User struct {
	ID         int
	Name       string
	ProfilePic int
	Followers  string
	Banned     string
	Photos     string
}

type Photo struct {
	ID       int
	UserId   int
	Path     string
	Likes    string
	Comments string
	Date     time.Time
}

type Comment struct {
	ID      int
	Content string
	PhotoId int
	UserId  int
	Date    time.Time
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// DoLogin  creates a new user in the database. It returns an updated User object (with the ID)
	DoLogin(User) (User, error)

	// getUserProfile gets the information of an user from its ID.
	GetUserProfile(User) (User, error)

	// getMyStream gets the stream of photos of the user searched from its ID
	GetMyStream(User) (User, error)

	// getLogo gets the profile picture of a user given its ID
	GetLogo(int) (int, error)

	// getImage gets a picture given its ID
	GetImage(Photo) (Photo, error)

	// uploadPhoto gets a path of an image and uploads the photo.
	UploadLogo(Photo, User) (Photo, User, error)

	// uploadLogo gets a path of an image and uploads the profile picture.
	UploadPhoto(Photo, User) (Photo, User, error)

	// commentPhoto inserts a comment on the comments table,
	CommentPhoto(Comment, Photo, User) (Comment, Photo, User, error)

	// likePhoto updates a photo and adds a like from a user,
	LikePhoto(Photo, User) (Photo, User, error)

	// setMyUserName updates a table with  comment on the comments table,
	SetMyUserName(User) (User, error)

	// followUser updates the list of followers of the user followed
	FollowUser(User, User) (User, error)

	// banUser updates the list of banned users of the user
	BanUser(User, User) (User, error)

	// unfollowUser removes the user followed from the list of followers of the user
	UnfollowUser(User, User) (User, error)

	// unbanUser removes the user banned from the list of banned users of the user1
	UnbanUser(User, User) (User, error)

	// unlike photo updates a photo and removes a like from an user,
	UnlikePhoto(Photo, User) (Photo, User, error)

	// uncomment photo updates a photo and removes a comment from an user,
	UncommentPhoto(Comment, Photo, User) (Comment, Photo, User, error)

	// deletePhoto removes a photo
	DeletePhoto(Photo, User) (Photo, User, error)

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	var err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		var sqlStmt = `CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
	profilepic INTEGER,
	followers TEXT NOT NULL,
	banned TEXT NOT NULL,
	photos TEXT NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err2 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err2, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    userid INTEGER NOT NULL,
	path TEXT NOT NULL,
	likes TEXT NOT NULL,
	comments TEXT NOT NULL,
	date DATE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err3 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err3, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
    commentid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	content TEXT NOT NULL,
    photoid INTEGER NOT NULL,
	userid INTEGER NOT NULL,
	date DATE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
