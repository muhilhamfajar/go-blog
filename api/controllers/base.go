package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-blog/api/models"
	"log"
	"net/http"
)

type Server struct {
	DB		*gorm.DB
	Router	*mux.Router
}

func (server *Server)Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string)  {
	var err error

	DBurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	server.DB, err = gorm.Open(Dbdriver, DBurl)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	}
	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string)  {
	fmt.Println("Listen to port 8081")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}