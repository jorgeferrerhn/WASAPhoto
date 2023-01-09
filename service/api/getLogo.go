package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

type ProfileLogo struct {
	logo uint64
}

func (rt *_router) getLogo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//this function receives a user id and returns the stream of photos of that user
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

	//Searchs for the user to get its logo

	logo, err := rt.db.GetLogo(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	//cast u.Photos to array
	//photos := strings.Fields(stringJson)

	//return the array on string format

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	a := `{"ProfileLogo":`
	a += strconv.FormatUint(logo, 10)
	a += "}"

	var p ProfileLogo

	err = json.Unmarshal([]byte(a), &p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	_ = json.NewEncoder(w).Encode(p)

}
