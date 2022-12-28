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
	json.Unmarshal(rowJson, &user)
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

	//función que recibe userId y devuelve el userProfile

	/*




		dbuser, err := rt.db.CreateUser(user.ToDatabase())

		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
		user.FromDatabase(dbuser)

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user)

	*/
}
