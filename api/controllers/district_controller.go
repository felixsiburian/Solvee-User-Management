package controllers

import (
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//get semua district
func (server *Server) GetDistricts(w http.ResponseWriter, r *http.Request) {
	district := models.Disctrict{}

	districts, err := district.FindAllDistrict(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, districts)
}

//get district berdasarkan ID
func(server *Server) GetDistrict (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	district := models.Disctrict{}

	districtReceived, err := district.FindDistrictByID(server.DB, pid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, districtReceived)
}

