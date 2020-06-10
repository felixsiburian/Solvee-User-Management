package controllers

import (
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//get semua regency
func (server *Server) GetRegencies (w http.ResponseWriter, r *http.Request){
	regenci := models.Regency{}

	regencies, err := regenci.FindALLRegency(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, regencies)
}

//get regenci berdasarkan id regency
func (server *Server) GetRegency(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	regenci := models.Regency{}

	regencyReceived, err := regenci.FindRegencyByID(server.DB, pid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, regencyReceived)
}


