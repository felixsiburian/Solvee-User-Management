package controllers

import (
	"Solvee-User-Management/api/auth"
	"Solvee-User-Management/api/models"
	"Solvee-User-Management/api/response"
	"Solvee-User-Management/utils/formaterror"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateReport(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	report := models.Report{}
	err = json.Unmarshal(body, &report)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	report.Prepare()
	err = report.Validate()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != report.UserID {
		response.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	reportCreated, err := report.SaveReport(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, reportCreated.ID))
	response.JSON(w, http.StatusCreated, reportCreated)
}

func (server *Server) GetReports(w http.ResponseWriter, r *http.Request) {
	report := models.Report{}

	reports, err := report.FindAllReport(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, reports)
}

func (server *Server) GetAReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	report := models.Report{}

	reportReceived, err := report.FindReportByID(server.DB, pid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, reportReceived)
}

func (server *Server) UpdateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//	cek laporan valid atau tidak
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	cek apakah tokennya valid dan kita akn mendapatkan id dari token tsb
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	cek apakah laporan tersebut eksis atau tidak
	report := models.Report{}
	err = server.DB.Debug().Model(models.Report{}).Where("id = ?", pid).Take(&report).Error
	if err != nil {
		response.ERROR(w, http.StatusNotFound, errors.New("Report Not Found"))
		return
	}

	//	jika ada user yang ingin mengupdate laporan yang bukan punya dia
	if uid != report.UserID {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	membaca data yg sudah di post
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//	 proses data
	reportUpdate := models.Report{}
	err = json.Unmarshal(body, &reportUpdate)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//	cek apakah user id request = user id yg didapat dari token
	if uid != reportUpdate.UserID {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	reportUpdate.Prepare()
	err = reportUpdate.Validate()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	reportUpdate.ID = report.ID

	reportUpdated, err := reportUpdate.UpdateAReport(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	response.JSON(w, http.StatusOK, reportUpdated)
}

func (server *Server) DeleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//	cek apakah laporan valid atau tidak
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	cek apakah user terautentikasi atau tidak
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//	cek laporan eksis atau tidak
	report := models.Report{}
	err = server.DB.Debug().Model(models.Report{}).Where("id = ?", pid).Take(&report).Error
	if err != nil {
		response.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	//	jika user telah terautentikasi, cek apakah laporan tsb dibuat oleh user ybs atau tidak
	if uid != report.UserID {
		response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = report.DeleteAReport(server.DB, pid, uid)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	response.JSON(w, http.StatusNoContent, "")
}
