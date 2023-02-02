package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLogo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// this function receives a user id and returns the stream of photos of that user

	id := ps.ByName("id")

	if id == "" {
		// ID not found
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, e := strconv.Atoi(id)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Searchs for the user to get its logo
	var p Photo
	p.ID = i

	var u User
	u.ProfilePic = i
	dbphoto, dbuser, err := rt.db.GetLogo(p.ToDatabase(), u.ToDatabase())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	//  Here we can re-use `photo` as FromDatabase is overwriting every variable in the structure.
	p.FromDatabase(dbphoto)
	u.FromDatabase(dbuser)

	//  Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)

	defer r.Body.Close()

}
