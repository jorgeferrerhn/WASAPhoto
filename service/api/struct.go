package api

import (
	"regexp"
	"time"

	"github.com/jorgeferrerhn/WASAPhoto/service/database"
)

// User struct represent a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See User.FromDatabase (below) to understand why.
type User struct {
	ID         uint64 `json:"Id"`
	Name       string `json:"Name"`
	ProfilePic uint64 `json:"ProfilePic"`
	Followers  string `json:"Followers"`
	Photos     string `json:"Photos"`
}

type Photo struct {
	ID       uint64    `json:"Id"`
	UserId   int       `json:"userId"`
	Path     string    `json:"path"`
	Likes    string    `json:"likes"`
	Comments string    `json:"comments"`
	Date     time.Time `json:"date"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.

func (u *User) FromDatabase(user database.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.ProfilePic = user.ProfilePic
	u.Followers = user.Followers
	u.Photos = user.Photos

}

// ToDatabase returns the user in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:         u.ID,
		Name:       u.Name,
		ProfilePic: u.ProfilePic,
		Followers:  u.Followers,
		Photos:     u.Photos,
	}
}

// FromDatabase populates the struct with data from the database, overwriting all values.

func (p *Photo) FromDatabase(photo database.Photo) {
	p.ID = photo.ID
	p.UserId = photo.UserId
	p.Path = photo.Path
	p.Likes = photo.Likes
	p.Comments = photo.Comments
	p.Date = photo.Date

}

// ToDatabase returns the user in a database-compatible representation
func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		ID:       p.ID,
		UserId:   p.UserId,
		Path:     p.Path,
		Likes:    p.Likes,
		Comments: p.Comments,
		Date:     p.Date,
	}
}

// IsValid checks the validity of the content.
func (u *User) IsValid() bool {
	m, err := regexp.MatchString("[a-zA-Z0-9]+", u.Name)

	if err != nil {
		return false
	}
	return m
}
