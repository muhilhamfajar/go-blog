package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, success bool, data interface{})  {
	var err error
	w.WriteHeader(statusCode)

	if !success {
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	err = json.NewEncoder(w).Encode(struct {
		Status	int	`json:"status"`
		Data	interface{} `json:"data"`
	}{
		Status: 1,
		Data: data,
	})

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error)  {
	if err != nil {
		JSON(w, statusCode, false, struct {
			Status	int		`json:"status"`
			Error 	string 	`json:"error"`
		}{
			Error: err.Error(),
			Status: 0,
		})
		return
	}

	JSON(w, http.StatusBadRequest, true, nil)
}