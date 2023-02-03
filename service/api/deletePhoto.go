package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Takes the photo Id and updates its like in the photos table
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

	// photo id
	photoId := ps.ByName("photoId")

	if photoId == "" {
		// Empty Photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intPhoto, err := strconv.Atoi(photoId)
	if err != nil {
		// id was not properly cast
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var p Photo
	var u User

	p.ID = intPhoto
	p.UserId = intId // only for sending it to the database function
	u.ID = intId

	// update info from database
	dbphoto, dbuser, err := rt.db.DeletePhoto(p.ToDatabase(), u.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't unlike the photo")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	// Here we can re-use `photo ` as FromDatabase is overwriting every variable in the structure.
	u.FromDatabase(dbuser)
	p.FromDatabase(dbphoto)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")

	if p.ID == 0 { // user not found
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Send the output to the user.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(u)

	defer r.Body.Close()

}
