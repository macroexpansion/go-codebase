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
	if err := c.db.Debug().Create(&account).Error; err != nil {
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

	var account models.Account
	if err := c.db.Debug().Where("username = ?", params.Username).First(&account).Error; err != nil {
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
