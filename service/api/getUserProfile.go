package api

import (
	"encoding/json"
	"errors"
	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func funcName(ps httprouter.Params) (int, error) {
	i := ps.ByName("id")

	if i == "" {
		//Empty ID
		return -1, errors.New("Empty ID")

	}

	intId, err := strconv.Atoi(i)
	if err != nil {
		// id wasn`t properly casted
		return -1, err

	}
	return intId, err
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	intId, err := funcName(ps)
	if err != nil {
		//error on database
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Searchs for the user to get the profile and returns the information
	var user User
	user.ID = intId
	dbuser, err := rt.db.GetUserProfile(user.ToDatabase())

	user.FromDatabase(dbuser)

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

}
