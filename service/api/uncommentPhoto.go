package api

import (
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
		//  id wasn`t properly casted
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
		//  id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create a Photo Struct
	var c Comment
	c.ID = intComment // default
	c.UserId = intId

	// update info from database
	res, err := rt.db.UncommentPhoto(c.ToDatabase())
	if err != nil {
		//  In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		//  Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}

	//  Send the output to the user.

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(res)))

	defer r.Body.Close()

}
