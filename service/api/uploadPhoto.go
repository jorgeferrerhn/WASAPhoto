package api

import (
	"bytes"
	"encoding/json"
	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Takes the userId and the path of the photo, and uploads it (updates the stream of photos)

	intId, err := checkId(ps)

	if err != nil {
		// error on database
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// path to the image
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r.Body)
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newStr := buf.String()

	// Here we have to cast the input content to a Struct
	// jsonPhoto := `{"id": 0, "userid": 0, "path": "", "likes": "", "comments": "", "date": ""}`
	var p Photo
	// err = json.Unmarshal([]byte(jsonPhoto), &p)

	p.UserId = intId
	p.Path = newStr

	/*
		p.UserId = intId
		p.Path = newStr


		u.ID = intId
	*/

	var u User
	u.ID = intId
	// update info from database
	dbphoto, dbuser, err := rt.db.UploadPhoto(p.ToDatabase(), u.ToDatabase())

	if err != nil {
		//  In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		//  Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't upload the photo")
		w.WriteHeader(http.StatusBadRequest) // 500
		return
	}
	//  Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
	p.FromDatabase(dbphoto)
	u.FromDatabase(dbuser)

	//  Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)

	defer r.Body.Close()

}
