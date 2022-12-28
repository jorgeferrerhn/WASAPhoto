package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var userId int
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

	//now, we will separate the userID and the imageID

	s := strings.Split(i, ".")
	a := s[0]
	b := s[1]

	userId, e1 := strconv.Atoi(a)
	imageId, e2 := strconv.Atoi(b)

	if e1 != nil || e2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	//imageId = strconv.Atoi(b)

	fmt.Println(userId)
	fmt.Println(imageId)

	//Searchs for the user to get its logo

	img, err := rt.db.GetImage(userId, imageId)
	fmt.Println(img)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "image/png")

}
