package controllers

import (
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (server *Server) GetVillages (w http.ResponseWriter, r *http.Request) {
	village := models.Village{}
	villages, err := village.FindAllVillage(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, villages)
}

func (server *Server) GetVillage (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	village := models.Village{}

	villageReceived, err := village.FindVillageByID(server.DB, pid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, villageReceived)
}
