package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gnsalok/go-microservices/mvc/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		//Bad request
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("UserId must be a no."))
		return

	}
	user, err := services.GetUser(userId)

	if err != nil {
		//Handle the err and return to client
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return

	}

	//return user to the client
	jsOnValue, _ := json.Marshal(user)
	res.Write(jsOnValue)
}
