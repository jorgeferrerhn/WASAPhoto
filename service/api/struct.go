package api

import "git.sapienzaapps.it/gamificationlab/wasa-fountains/tree/master/service/database" //este import da error

// User struct represent a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See User.FromDatabase (below) to understand why.
type User struct {
	Name            string   `json:"name"`
	Profile_picture string   `json:"profilepic"`
	PhotosStream    []string `json:"photos"`
	//falta el resto de atributos
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.User" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *User) FromDatabase(user database.User) {
	u.Name = user.Name
	u.Profile_picture = user.Profile_picture
	u.PhotosStream = user.PhotosStream
	//falta el resto de atributos

}

// ToDatabase returns the fountain in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		Name:            u.Name,
		Profile_picture: u.Profile_picture,
		PhotosStream:    u.PhotosStream,
	}
}
