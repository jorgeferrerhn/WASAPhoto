package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var imageId int

	//This function receives an image ID and returns the image.
	//Due to the complexity of the search of an image ID through a SQL Text (Photos), we will suppose that the image ID is a string with the format <userID>.<imageID>
	//For example, to get the third image posted from the user ID 3, we will get the image ID "3.3", then the image ID will have string type.
	r.Close = true

	i := ps.ByName("id")

	if i == "" {
		//Empty ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	imageId, e2 := strconv.Atoi(i)

	if e2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	fmt.Println(imageId)

	//Searchs for the user to get its logo

	var p Photo

	p.ID = imageId
	dbphoto, err := rt.db.GetImage(p.ToDatabase())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	// Here we can re-use `photo` as FromDatabase is overwriting every variabile in the structure.
	p.FromDatabase(dbphoto)

	if p.Path == "" { //photo didn't exist
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)

	defer r.Body.Close()

}
