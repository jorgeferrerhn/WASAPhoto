package api

import (
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLogo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//this function receives a user id and returns the stream of photos of that user

	id := ps.ByName("id")

	if id == "" {
		//ID not found
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, e := strconv.Atoi(id)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Searchs for the user to get its logo

	logo, err := rt.db.GetLogo(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	// Send the output to the user.

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(logo)))

	defer r.Body.Close()

}
