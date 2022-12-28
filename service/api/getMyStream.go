package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	var user User
	stringJson, err := rt.db.GetMyStream(i)
	fmt.Println(stringJson)

	//cast u.Photos to array
	//photos := strings.Fields(stringJson)

	fmt.Println("NOW AN EXAMPLE")
	//example
	example := `
				{"PhotoId": 100000, 
                          "Likes": ["JohnDoe","CarlSagan"],
                          "Comments": [
                              {"User":"JohnDoe","Comment":"Wow, nice trip!"},
                              {"User":"CarlSagan", "Comment":"Amazing!"}],
                          "Date":"2021-01-30T08:30:00Z",
                          "Followers": ["JohnDoe","KenFollet"]
                 },
				 {"PhotoId": 100001, 
                          "Likes": ["JohnDoe","CarlSagan"],
                          "Comments": [
                              {"User":"JohnDoe","Comment":"Wow, nice trip!"},
                              {"User":"CarlSagan", "Comment":"Amazing!"}],
                          "Date":"2021-01-30T08:30:00Z",
                          "Followers": ["JohnDoe","KenFollet"]
                 }
				`

	photos := strings.Fields(example)

	fmt.Println(photos)

	//cast string to JSON

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

}
