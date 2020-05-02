package controllers

import (
	"go-blog/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request)  {
	responses.JSON(w, http.StatusOK, true,"API GO BLOG")
}
