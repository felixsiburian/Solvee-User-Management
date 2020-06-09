package controllers

import (
	"Solvee-User-Management/api/response"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}