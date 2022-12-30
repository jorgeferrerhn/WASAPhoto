package api

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

type Path struct {
	path string
}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//Takes the userId and the path of the photo, and uploads it (updates the stream of photos)

	//user id
	i := ps.ByName("id")
	fmt.Println("ID: ", i)

	if i == "" {
		//Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(i)
	if err != nil {
		// id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//path to the image
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()

	fmt.Println("Path: ", newStr)

	//update info from database
	id, err := rt.db.UploadPhoto(intId, newStr)
	fmt.Println(id)
	if err != nil {
		// error updating database
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if !user.IsValid() {
			// Here we validated the user structure content (correct name), and we
			// discovered that the user data are not valid.
			// Note: the IsValid() function skips the ID check (see below).
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dbuser, err := rt.db.UploadPhoto()

		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
		user.FromDatabase(dbuser)

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user)

		defer r.Body.Close()
	*/
}
