package api

import (
	"net/http"
	"strconv"

	"github.com/jorgeferrerhn/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	//Searchs for the user to get the stream of photos

	stringJson, err := rt.db.GetMyStream(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	//cast u.Photos to array
	//photos := strings.Fields(stringJson)

	stringJson = `{"PhotoId": 100000,"Likes": ["JohnDoe","CarlSagan"],"Comments": [{"User":"JohnDoe","Comment":"Wow, nice trip!"},{"User":"CarlSagan", "Comment":"Amazing!"}],"Date":"2021-01-30T08:30:00Z","Followers": ["JohnDoe","KenFollet"]};{"PhotoId": 100001,"Likes": ["JohnDoe","CarlSagan"],"Comments": [{"User":"JohnDoe","Comment":"Wow, nice trip!"},{"User":"CarlSagan", "Comment":"Amazing!"}],"Date":"2021-01-30T08:30:00Z","Followers": ["JohnDoe","KenFollet"]}`
	//faltaría decodificar la información en cada uno de los structs (comentarios, fechas, etc)
	/*example
	example := `
		{"PhotoId": 100000,
		"Likes": ["JohnDoe","CarlSagan"],
		"Comments": [
			{"User":"JohnDoe","Comment":"Wow, nice trip!"},
			{"User":"CarlSagan", "Comment":"Amazing!"}],
		"Date":"2021-01-30T08:30:00Z",
		"Followers": ["JohnDoe","KenFollet"]
			};
			{"PhotoId": 100001,
		"Likes": ["JohnDoe","CarlSagan"],
		"Comments": [
			{"User":"JohnDoe","Comment":"Wow, nice trip!"},
			{"User":"CarlSagan", "Comment":"Amazing!"}],
		"Date":"2021-01-30T08:30:00Z",
		"Followers": ["JohnDoe","KenFollet"]
			}
		`
	*/

	//return the array on string format

	defer r.Body.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(stringJson))

}
