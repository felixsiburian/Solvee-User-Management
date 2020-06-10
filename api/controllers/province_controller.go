package controllers

import (
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (server *Server) GetProvinces (w http.ResponseWriter, r *http.Request) {
	province := models.Province{}

	provinces, err := province.FindAllProvince(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, provinces)
}

func (server *Server) GetProvince(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	province := models.Province{}

	provinceReceived, err := province.FindProvinceByID(server.DB, pid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, provinceReceived)
}
