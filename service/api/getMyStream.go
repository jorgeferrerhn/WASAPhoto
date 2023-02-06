package api

import (
	"encoding/json"
	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// This function receives a user id and returns the stream of photos of that user

	intId, err := checkId(ps)
	if err != nil {
		// error on database
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Search for the user to get the stream of photos
	var user User
	user.ID = intId

	dbuser, err2 := rt.db.GetMyStream(user.ToDatabase())
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	user.FromDatabase(dbuser)

	in := []byte(user.Photos)
	var castPhotos []Photo
	err3 := json.Unmarshal(in, &castPhotos)
	if err3 != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	err4 := json.NewEncoder(w).Encode(castPhotos)
	if err4 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	defer r.Body.Close()

}
