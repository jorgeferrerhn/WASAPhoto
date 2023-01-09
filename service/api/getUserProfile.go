package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	r.Close = true

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

	//Searchs for the user to get the profile and returns the information

	var user User
	rowJson, err := rt.db.GetUserProfile(i)
	err = json.Unmarshal(rowJson, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	//cast to string

	if user.ID == 0 && user.Name == "" {
		//user not found
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

}
