package controllers

import "github.com/muhilhamfajar/go-blog/api/middlewares"

func (server *Server) InitializeRoutes()  {
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJson(server.Home)).Methods("GET")

	// Login Routes
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJson(server.Login)).Methods("POST")

	// User Routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(server.GetUser)).Methods("GET")
}
