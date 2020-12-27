package controllers

import (
	"net/http"
)

type userController struct {
}

func (uc userController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}

func newUserController() *userController {
	return &userController{}
}
