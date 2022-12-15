package api

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//función que recibe username. Busca en la base de datos, si no existe lo crea.
	//Devuelve identificador del usuario

	body, err := io.ReadAll(r.Body)

	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := string(body)

	//comprobación de si el usuario existe
	fmt.Println(username)

	//comprobamos que el usuario coincida con el Regex
	m, err := regexp.MatchString("[a-zA-Z0-9]+", username)
	fmt.Println(m)

	//casteamos al user

	//buscamos el usuario en la base de datos. Si existe, devolvemos su ID. Si no, lo creamos y generamos un nuevo ID:

	//aquí iría la parte de búsqueda en la base de datos

	//este código se ejecuta en caso de que no se encuentre en la base de datos
	var user User
	user.Name = username
	user.Id = rand.Intn(10000)
	fmt.Println(user)

	// Create the user in the database. Note that this function will return a new instance of the user with the
	// same information, plus the ID.
	dbuser, err := rt.db.doLogin(user.ToDatabase()) //esta función falta implementarla en el database
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the fountain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
	user.FromDatabase(dbuser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Hi %d!", user.Id)

	//_ = json.NewEncoder(w).Encode(user)
}
