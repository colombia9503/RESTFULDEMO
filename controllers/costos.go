package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-SqlServer/common"
	"github.com/colombia9503/RESTful-SqlServer/models"
	"github.com/gorilla/mux"
)

type costosController struct{}

var Costos = new(costosController)

func (cc *costosController) Get(w http.ResponseWriter, r *http.Request) {
	costs, err := models.Costos.SelectAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
	}

	res, err := json.Marshal(costs)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (cc *costosController) Exists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := models.Costos.Exists(id)
	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (cc *costosController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := models.Costos.SelectOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}
	res, err := json.Marshal(result)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}

func (cc *costosController) Insert(w http.ResponseWriter, r *http.Request) {
	var cost models.Costo
	if err := json.NewDecoder(r.Body).Decode(&cost); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	verif, err := models.Costos.Insert(cost)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}
	res, err := json.Marshal(verif)
	if err != nil {
		common.JsonError(w, err, http.StatusNotFound)
		return
	}

	common.JsonOk(w, res, http.StatusOK)
}
