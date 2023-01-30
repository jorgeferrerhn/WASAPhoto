package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//  Takes the photo Id and updates its like in the photos table
	//  user id
	i := ps.ByName("id")

	if i == "" {
		// Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(i)
	if err != nil {
		//  id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  followedUser
	id_followed := ps.ByName("id2")

	if id_followed == "" {
		// Empty Followed ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intFollowed, err := strconv.Atoi(id_followed)
	if err != nil {
		//  id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// User 1
	var u1 User
	u1.ID = intId

	// User 2
	var u2 User
	u2.ID = intFollowed

	// update info from database
	dbuser1, err := rt.db.BanUser(u1.ToDatabase(), u2.ToDatabase())
	if err != nil {
		//  In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		//  Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't update the banned list")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	u1.FromDatabase(dbuser1)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(u1)

	defer r.Body.Close()

}
