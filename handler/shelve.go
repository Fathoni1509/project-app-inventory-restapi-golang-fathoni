package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type ShelveHandler struct {
	ShelveService service.ShelveService
	Config utils.Configuration
}

func NewShelveHandler(shelveService service.ShelveService, config utils.Configuration) ShelveHandler {
	return ShelveHandler{
		ShelveService: shelveService,
		Config: config,
	}
}

// create shelve
func (shelveHandler *ShelveHandler) AddShelve(w http.ResponseWriter, r *http.Request) {
	var req dto.ShelveCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data", nil)
		return
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model shelve
	shelve := model.Shelve{
		WarehouseId: req.WarehouseId,
		Name: req.Name,
	}

	// create shelve service
	err = shelveHandler.ShelveService.AddShelve(&shelve)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created shelve", nil)
}

// get list shelves
func (shelveHandler *ShelveHandler) GetListShelves(w http.ResponseWriter, r *http.Request) {
	// Get data shelve form service all shelve
	shelves, err := shelveHandler.ShelveService.GetListShelves()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch shelves: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data shelve", shelves)

}

// get list shelve by id
func (shelveHandler *ShelveHandler) GetListShelveByID(w http.ResponseWriter, r *http.Request) {
	shelveIDstr := chi.URLParam(r, "shelve_id")

	shelveID, err := strconv.Atoi(shelveIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid shelve ID", nil)
		return
	}

	response, err := shelveHandler.ShelveService.GetListShelveById(shelveID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "shelve not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data shelve by id", response)
}

// update shelve
func (shelveHandler *ShelveHandler) UpdateShelve(w http.ResponseWriter, r *http.Request) {
	shelveIDstr := chi.URLParam(r, "shelve_id")

	shelveID, err := strconv.Atoi(shelveIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param shelve id :"+err.Error(), nil)
		return
	}

	var req dto.ShelveUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := shelveHandler.ShelveService.GetListShelveById(shelveID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.WarehouseId != nil {
		existing.WarehouseId = *req.WarehouseId
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model shelve
	shelve := model.Shelve{
		WarehouseId: existing.WarehouseId,
		Name: existing.Name,
	}

	err = shelveHandler.ShelveService.UpdateShelve(shelveID, &shelve)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete shelve
func (shelveHandler *ShelveHandler) DeleteShelve(w http.ResponseWriter, r *http.Request) {
	shelveIDstr := chi.URLParam(r, "shelve_id")

	shelveID, err := strconv.Atoi(shelveIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid shelve ID", nil)
		return
	}

	err = shelveHandler.ShelveService.DeleteShelve(shelveID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}