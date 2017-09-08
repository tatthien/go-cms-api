package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"msg": "hello",
	}
	sendJSON(w, http.StatusOK, data)
}

func (s *Server) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.GetPosts(0, 10)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, err.Error())
	}
	sendJSON(w, http.StatusOK, posts)
}

func (s *Server) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		sendJSON(w, http.StatusBadRequest, "Invalid post id")
		return
	}

	post, err := s.db.GetPost(id)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	sendJSON(w, http.StatusOK, post)
	return
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "private; max-age=86400")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err.Error())
	}
}
