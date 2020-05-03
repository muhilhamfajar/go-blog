package controllers

import (
	"encoding/json"
	"github.com/muhilhamfajar/go-blog/api/auth"
	"github.com/muhilhamfajar/go-blog/api/models"
	"github.com/muhilhamfajar/go-blog/api/responses"
	"github.com/muhilhamfajar/go-blog/api/utils/formaterror"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if	err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if	err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, true, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	user := models.User{}

	getUser, err := user.GetUserByEmail(server.DB, email)
	if err != nil {
		return "", err
	}

	err = user.ValidatePassword(getUser.Password, password)
	if err != nil {
		return "", err
	}

	return auth.CreateToken(getUser.ID)
}
