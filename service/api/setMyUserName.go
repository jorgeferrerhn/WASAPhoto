package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	reqToken := r.Header.Get("content-type")
	token, errTok := strconv.Atoi(reqToken)
	if errTok != nil {
		// id was not properly cast
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Takes the userId and the path of the photo, and uploads it (updates the stream of photos)

	// user id
	i := ps.ByName("id")

	if i == "" {
		// Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(i)
	if err != nil {
		// id was not properly cast
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// new user name
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r.Body)
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := buf.String()

	// create user
	var user User
	user.ID = intId
	user.Name = name // updating the name here

	if !user.IsValid() {
		// Here we validated the user structure content (correct name), and we discovered that the user data is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.ID != token {
		// Error: the authorization header is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update info from database
	dbuser, err := rt.db.SetMyUserName(user.ToDatabase())

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't update the username")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	// Here we can re-use `user` as FromDatabase is overwriting every variable in the structure.
	user.FromDatabase(dbuser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

	defer r.Body.Close()
}
