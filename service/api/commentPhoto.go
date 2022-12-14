package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//Takes the userId and the comment, and uploads it (updates the comments table)

	// user id
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
	fmt.Println(intId)

	// user id
	photoId := ps.ByName("photoId")
	fmt.Println("Photo ID: ", photoId)

	if photoId == "" {
		//Empty Photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intPhoto, err := strconv.ParseUint(photoId, 10, 64)
	if err != nil {
		// id wasn`t properly casted
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(intPhoto)

	//Comment
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	comment := buf.String()

	fmt.Println("Comment: ", comment)

	//create a Photo Struct
	var c Comment

	c.UserId = intId
	c.Content = comment
	c.PhotoId = intPhoto

	//update info from database
	dbcomment, err := rt.db.CommentPhoto(c.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError) //500
		return
	}

	// Here we can re-use `comment` as FromDatabase is overwriting every variabile in the structure.
	c.FromDatabase(dbcomment)

	if c.ID == 0 { //user not found
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(c)

	defer r.Body.Close()

}
