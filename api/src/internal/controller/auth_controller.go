package controller

import "net/http"
import "encoding/json"
import "gorm.io/gorm"
import "pgsql/models"

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db: db}
}

type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	if err := c.db.Create(&account).Error; err != nil {
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

	var account models.Account
	if err := c.db.Where(&models.Account{Username: params.Username}).First(&account).Error; err != nil {
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
