package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/tatthien/go-cms-api/database"
)

type Server struct {
	ip        string
	port      string
	db        *database.Database
	signKey   []byte
	verifyKey []byte
}

type Route struct {
	Path         string
	Method       string
	Handler      http.HandlerFunc
	Authenticate bool
}

type Routes []Route

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

const (
	privKeyPath = "/keys/app.rsa"
	pubKeyPath  = "/keys/app.rsa.pub"
)

var VerifyKey, SignKey []byte

func initKeys() {
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

// New start new server with ip and port
func New(ip, port string) *Server {
	db := database.Connect()
	return &Server{
		ip,
		port,
		db,
		SignKey,
		VerifyKey,
	}
}

// Run start listening the routes
func (s *Server) Run() {
	// Init keys

	r := mux.NewRouter()

	// Create routes
	routes := Routes{
		Route{"/api/v1", "GET", s.IndexHandler, false},
		Route{"/api/v1/check-login", "GET", s.CheckLoginHandler, true},
		Route{"/api/v1/login", "POST", s.LoginHandler, false},
		Route{"/api/v1/posts", "GET", s.GetPostsHandler, false},
		Route{"/api/v1/posts", "POST", s.StorePostHandler, true},
		Route{"/api/v1/posts/{slug}", "GET", s.GetPostHandler, false},
		Route{"/api/v1/posts/{id}", "PATCH", s.UpdatePostHandler, true},
		Route{"/api/v1/posts/{id}", "DELETE", s.DeletePostHandler, true},
	}
	for _, route := range routes {
		if route.Authenticate {
			r.Handle(route.Path, negroni.New(
				negroni.HandlerFunc(s.ValidateTokenMiddleware),
				negroni.Wrap(route.Handler),
			))
		} else {
			r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	fmt.Printf("Sever is started at %s:%s\n", s.ip, s.port)
	log.Fatal(http.ListenAndServe(s.ip+":"+s.port, r))
}

// Close the server
func (s *Server) Close() {
	defer s.db.Close()
}
