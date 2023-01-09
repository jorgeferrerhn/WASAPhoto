package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	//photo id
	p := ps.ByName("photoId")
	fmt.Println("Photo ID: ", p)

	if p == "" {
		//Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intPhotoId, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		// id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//create a Photo Struct
	var p2 Photo
	p2.ID = intPhotoId //default
	p2.UserId = intId

	//update info from database
	res, err := rt.db.DeletePhoto(p2.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError) //500
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(res)))

	defer r.Body.Close()

}
