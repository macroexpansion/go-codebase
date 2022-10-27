package controller

import (
	"encoding/json"
	"net/http"

	"pgsql/models"
	"pgsql/repos"
)

type AuthController struct {
	acc repos.AccountInterface
}

func NewAuthController(acc repos.AccountInterface) *AuthController {
	return &AuthController{acc: acc}
}

type AuthParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *AuthController) Register(res http.ResponseWriter, req *http.Request) {
	var params AuthParam
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(res).Encode(&Response{Message: "Register Error"})
		return
	}

	account := models.Account{
		Username: params.Username,
		Password: params.Password,
	}
	if err := c.acc.Save(&account); err != nil {
		json.NewEncoder(res).Encode(&Response{Message: err.Error()})
		return
	}

	json.NewEncoder(res).Encode(&Response{
		Status:  "ok",
		Message: "registered",
		Data:    map[string]string{"username": params.Username},
	})
}

func (c *AuthController) Login(res http.ResponseWriter, req *http.Request) {
	var params AuthParam
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(res).Encode(&Response{Message: "Login Error"})
		return
	}

	validation := validate.Struct(params)
	if validation != nil {
		json.NewEncoder(res).Encode(&Response{Message: "Validate Error"})
		return
	}

	account, err := c.acc.FindByUsername(params.Username)
	if err != nil {
		json.NewEncoder(res).Encode(&Response{Message: "Login Error"})
		return
	}

	if account.Password != params.Password {
		json.NewEncoder(res).Encode(&Response{Message: "Login Error"})
		return
	}

	json.NewEncoder(res).Encode(&Response{
		Status:  "ok",
		Message: "logged in",
		Data:    map[string]string{"token": "TOKENQWE123"},
	})
}
