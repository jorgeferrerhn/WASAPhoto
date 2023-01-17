package api

import (
	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//This function receives a user id and returns the stream of photos of that user

	intId, err := checkId(ps)
	if err != nil {
		//error on database
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Searchs for the user to get the stream of photos

	stringJson, err := rt.db.GetMyStream(intId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(stringJson))

}
