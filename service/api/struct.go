package api

import (
	"regexp"

	"github.com/jorgeferrerhn/WASAPhoto/service/database"
) //este import da error

// User struct represent a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See User.FromDatabase (below) to understand why.
type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.

func (u *User) FromDatabase(user database.User) {
	u.ID = user.ID
	u.Name = user.Name

}

// ToDatabase returns the user in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:   u.ID,
		Name: u.Name,
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
