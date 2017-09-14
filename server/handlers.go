package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/tatthien/go-cms-api/model"
)

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "private; max-age=86400")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err.Error())
	}
}

// ValidateTokenMiddleware validate token middleware
func (s *Server) ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return s.verifyKey, nil
	})

	var resp Response

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			resp.Message = "Token is not valid"
			sendJSON(w, http.StatusUnauthorized, resp)
		}
	} else {
		resp.Message = "Unauthorized access to this resource"
		sendJSON(w, http.StatusUnauthorized, resp)
	}
}

// IndexHandler handler function for `/api/v1` endpoint
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		true,
		"Hello",
		nil,
	}
	sendJSON(w, http.StatusOK, resp)
}

// LoginHandler handler function for `/api/v1/login` endpoint
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser model.UserCredentials
	var resp Response
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		resp.Message = "Error in request"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	if loginUser.Email == "" {
		resp.Message = "Email is missing"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	if loginUser.Password == "" {
		resp.Message = "Password is missing"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	storedUser, err := s.db.GetUserByEmail(loginUser.Email)
	if err != nil {
		resp.Message = "Email does not match"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginUser.Password)); err != nil {
		resp.Message = "Password does not match"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	// Create token & send to user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(s.signKey)
	if err != nil {
		resp.Message = err.Error()
		sendJSON(w, http.StatusOK, resp)
		return
	}

	resp.Success = true
	resp.Data = map[string]interface{}{
		"token": tokenString,
		"user":  storedUser,
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

	resp := Response{false, "", nil}
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
	resp := Response{false, "", nil}

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

// StorePostHandler handler function for `/api/v1/posts/` endpoint
// store post data into database
func (s *Server) StorePostHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	var resp Response
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		resp.Message = "Error in request"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	post, err := s.db.InsertPost(post)
	if err != nil {
		resp.Message = err.Error()
		sendJSON(w, http.StatusOK, resp)
		return
	}

	resp.Success = true
	resp.Data = post
	sendJSON(w, http.StatusOK, resp)
}

// UpdatePostHandler handler function to update post endpoint
func (s *Server) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var resp Response
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		resp.Message = "Post ID is not valid"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		resp.Message = "Error in request"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	post, err = s.db.UpdatePost(id, post)
	resp.Success = true
	resp.Data = post

	sendJSON(w, http.StatusOK, resp)
}

// DeletePostHandler handler function to delete post
func (s *Server) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var resp Response
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		resp.Message = "Post ID is not valid"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	err = s.db.DeletePost(id)
	if err != nil {
		resp.Message = "Can not delete post"
		sendJSON(w, http.StatusOK, resp)
		return
	}

	resp.Success = true
	sendJSON(w, http.StatusOK, resp)
}
