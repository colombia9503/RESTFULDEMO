package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-SqlServer/common"
	"github.com/colombia9503/RESTful-SqlServer/models"
)

var Auth = new(authController)

type authController struct{}

func (ac *authController) Login(w http.ResponseWriter, r *http.Request) {
	var token string
	var u models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		common.JsonError(w, err, http.StatusInternalServerError)
		return
	}

	result, err := models.Auth.Login(u)
	if err != nil {
		common.JsonError(w, err, http.StatusUnauthorized)
		return
	}

	token, err = common.GenerateJWT(u.Nombre, u.Usuario)
	if err != nil {
		common.JsonError(w, err, http.StatusInternalServerError)
	}

	result.Token = token
	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}
