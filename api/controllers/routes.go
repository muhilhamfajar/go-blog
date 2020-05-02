package controllers

import "go-blog/api/middlewares"

func (server *Server) InitializeRoutes()  {
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJson(server.Home)).Methods("GET")

	// User Routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(server.CreateUser)).Methods("POST")
}
