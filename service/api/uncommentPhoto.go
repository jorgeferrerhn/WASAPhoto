package api

import (
	"encoding/json"
	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Takes the userId and the comment, and uploads it (updates the comments table)

	//  user id
	i := ps.ByName("id")

	if i == "" {
		// Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(i)
	if err != nil {
		//  id was not properly cast
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  comment id
	commentId := ps.ByName("commentId")

	if commentId == "" {
		// Empty Photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intComment, err := strconv.Atoi(commentId)
	if err != nil {
		//  id was not properly cast
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create a Photo Struct
	var c Comment
	c.ID = intComment // default
	c.UserId = intId

	// We pass this photo because we will have to update its comments in the database
	var p Photo

	var u User
	u.ID = intId

	// update info from database
	dbcomment, dbphoto, dbuser, err := rt.db.UncommentPhoto(c.ToDatabase(), p.ToDatabase(), u.ToDatabase())
	if err != nil {
		//  In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		//  Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't delete the comment")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//  Here we can re-use `comment` as FromDatabase is overwriting every variable in the structure.
	c.FromDatabase(dbcomment)
	p.FromDatabase(dbphoto)
	u.FromDatabase(dbuser)

	//  Send the output to the user.
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p) // Return the photo

	defer r.Body.Close()

}
