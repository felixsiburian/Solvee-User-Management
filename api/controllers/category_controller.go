package controllers

import (
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"net/http"
)

//get semua category
func (server *Server) GetCategory (w http.ResponseWriter, r *http.Request) {
	category := models.Category{}

	categories, err := category.FindAllCategories(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, categories)
}
