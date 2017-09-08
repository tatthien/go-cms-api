package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tatthien/go-cms-api/database"
)

type Server struct {
	ip   string
	port string
	db   *database.Database
}

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// New start new server with ip and port
func New(ip, port string) *Server {
	db := database.Connect()
	return &Server{
		ip,
		port,
		db,
	}
}

// Run start listening the routes
func (s *Server) Run() {
	r := mux.NewRouter()

	// Create routes
	routes := Routes{
		Route{"/api/v1", "GET", s.IndexHandler},
		Route{"/api/v1/posts", "GET", s.GetPostsHandler},
		Route{"/api/v1/posts/{id}", "GET", s.GetPostHandler},
	}
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	fmt.Printf("Sever is started at %s:%s\n", s.ip, s.port)
	log.Fatal(http.ListenAndServe(s.ip+":"+s.port, r))
}

// Close the server
func (s *Server) Close() {
	defer s.db.Close()
}
