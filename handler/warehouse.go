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

type WarehouseHandler struct {
	WarehouseService service.WarehouseService
	Config utils.Configuration
}

func NewWarehouseHandler(warehouseService service.WarehouseService, config utils.Configuration) WarehouseHandler {
	return WarehouseHandler{
		WarehouseService: warehouseService,
		Config: config,
	}
}

// create warehouse
func (warehouseHandler *WarehouseHandler) AddWarehouse(w http.ResponseWriter, r *http.Request) {
	var req dto.WarehouseCreateRequest
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

	// parsing to model warehouse
	warehouse := model.Warehouse{
		Name: req.Name,
		Location: req.Location,
	}

	// create category service
	err = warehouseHandler.WarehouseService.AddWarehouse(&warehouse)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created warehouse", nil)
}

// get list warehouses
func (warehouseHandler *WarehouseHandler) GetListWarehouses(w http.ResponseWriter, r *http.Request) {
	// Get data warehouse form service all warehouse
	warehouses, err := warehouseHandler.WarehouseService.GetListWarehouses()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch categories: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data warehouse", warehouses)

}

// get list warehouse by id
func (warehouseHandler *WarehouseHandler) GetListWarehouseByID(w http.ResponseWriter, r *http.Request) {
	warehouseIDstr := chi.URLParam(r, "warehouse_id")

	warehouseID, err := strconv.Atoi(warehouseIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid warehouse ID", nil)
		return
	}

	response, err := warehouseHandler.WarehouseService.GetListWarehouseById(warehouseID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Warehouse not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data warehouse by id", response)
}

// update warehouse
func (warehouseHandler *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouseIDstr := chi.URLParam(r, "warehouse_id")

	warehouseID, err := strconv.Atoi(warehouseIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param warehouse id :"+err.Error(), nil)
		return
	}

	var req dto.WarehouseUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := warehouseHandler.WarehouseService.GetListWarehouseById(warehouseID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.Location != nil {
		existing.Location = *req.Location
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model warehouse
	warehouse := model.Warehouse{
		Name: existing.Name,
		Location: existing.Location,
	}

	err = warehouseHandler.WarehouseService.UpdateWarehouse(warehouseID, &warehouse)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete warehouse
func (warehouseHandler *WarehouseHandler) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouseIDstr := chi.URLParam(r, "warehouse_id")

	warehouseID, err := strconv.Atoi(warehouseIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid warehouse ID", nil)
		return
	}

	err = warehouseHandler.WarehouseService.DeleteWarehouse(warehouseID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}