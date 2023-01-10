package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//función que recibe username. Busca en la base de datos, si no existe lo crea.
	//Devuelve identificador del usuario
	var user User

	fmt.Println(r.Body)
	//err := json.NewDecoder(r.Body).Decode(&user)

	//new user name
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	user.Name = buf.String()

	if !user.IsValid() {
		// Here we validated the user structure content (correct name), and we
		// discovered that the user data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("USER before: ", user.Name)

	dbuser, err := rt.db.DoLogin(user.ToDatabase())

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
	user.FromDatabase(dbuser)

	fmt.Println(user)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

	defer r.Body.Close()

}
