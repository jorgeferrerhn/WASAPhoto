package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	//new user name
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	name := buf.String()

	fmt.Println("New user name: ", name)

	//update info from database
	ret, err := rt.db.SetMyUserName(intId, name)
	fmt.Println(ret)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError) //500
		return
	}

	a := `{"Result":`
	a += strconv.Itoa(ret)
	a += "}"
	_ = json.NewEncoder(w).Encode(a)

	defer r.Body.Close()

}
