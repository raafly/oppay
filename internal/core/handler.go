package core

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/ewallet/pkg/helper"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type userHandler struct {
	port UserService
}

func NewUserHandler(Port UserService) UserHandler {
	return &userHandler{port: Port}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := userSignUp{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		panic(err)
	}

	if err = h.port.signUp(request); err != nil {
		panic(err)
	} else {
		response := webResponse {
			Code: 201,
			Status: "OK",
		}
		helper.WriteToResponBody(w, response)
	}
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := userSignIn{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		panic(err)
	}

	if err = h.port.singIn(request); err != nil {
		panic(err)
	} else {
		response := webResponse {
			Code: 201,
			Status: "OK",
		}
		helper.WriteToResponBody(w, response)
	}
}

func (h userHandler) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	if user, err := h.port.findById(userId); err != nil {
		panic(NewNotFoundError(err.Error()))
	} else {
		response := webResponse {
			Code: 200,
			Status: "OK",
			Data: user,
		}

		helper.WriteToResponBody(w, response)
	}
}