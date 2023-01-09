package api

import (
	"net/http"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	/*
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

		//create a Photo Struct
		var p Photo
		p.ID = 0 //default
		p.UserId = intId
		p.Path = newStr
		p.Comments = "[]"
		p.Date = time.Now()
		p.Likes = "[]"

		//update info from database
		dbphoto, err := rt.db.CommentPhoto(p.ToDatabase())
		fmt.Println(dbphoto)
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't upload the photo")
			w.WriteHeader(http.StatusInternalServerError) //500
			return
		}
		// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
		p.FromDatabase(dbphoto)

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(p)

		defer r.Body.Close()

	*/

}
