package handlers

import (
	// "context"
	"encoding/json"
	// "github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	// "io"
	"github.com/DmitriyKalekin/stalker22/dto"
	"net/http"
)

// RespondwithJSON write json response format with application/json header
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Request Handler - GET /posts - Read a list of posts.
// ListPosts godoc
// @Summary      ListPosts responds with the list of all albums as JSON.
// @Description  ListPosts responds with the list of all albums as JSON.
// @Tags         albums
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Album
// @Failure 401 {object} dto.HTTPError
// @Router       /posts [get]
func ListPosts(w http.ResponseWriter, r *http.Request) {
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// defer resp.Body.Close()

	// w.Header().Set("Content-Type", "application/json")

	// if _, err := io.Copy(w, resp.Body); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	var rrr = dto.Album{
		ID:     "5",
		Title:  "OK",
		Artist: "OK",
		Price:  3.14,
	}
	RespondwithJSON(w, http.StatusOK, rrr)
}

type Response struct {
	Status string `json:"status"`
}

func TgHandler(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	json.NewDecoder(r.Body).Decode(&res)
	log.Warnf("%#v", res)
	var rrr = Response{Status: "OK"}
	RespondwithJSON(w, http.StatusOK, rrr)
}
