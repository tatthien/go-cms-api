package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// IndexHandler handler function for `/api/v1` endpoint
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		true,
		"Hello",
		nil,
	}
	sendJSON(w, http.StatusOK, resp)
}

// GetPostsHandler handler function for `/api/v1/posts` endpoint
func (s *Server) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var args = map[string]string{}

	values := r.URL.Query()

	// Get limit from request params
	limit, ok := values["limit"]
	if ok && len(limit) > 0 {
		args["limit"] = limit[0]
	}

	// Get page from request params
	page, ok := values["page"]
	if ok && len(page) > 0 {
		args["page"] = page[0]
	}

	// Get post type from request params
	postType, ok := values["post_type"]
	if ok && len(postType) > 0 {
		args["post_type"] = postType[0]
	}

	posts, err := s.db.GetPosts(args)

	resp := Response{
		false,
		"",
		nil,
	}
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Data = posts
	}

	sendJSON(w, http.StatusOK, resp)
	return
}

// GetPostHandler handler function for `/api/v1/posts/{slug}` endpoint
func (s *Server) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	resp := Response{
		false,
		"",
		nil,
	}

	post, err := s.db.GetPostBySlug(slug)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Data = post
	}
	sendJSON(w, http.StatusOK, resp)
	return
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "private; max-age=86400")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err.Error())
	}
}
