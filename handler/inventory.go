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

type InventoryHandler struct {
	InventoryService service.InventoryService
	Config utils.Configuration
}

func NewInventoryHandler(inventoryService service.InventoryService, config utils.Configuration) InventoryHandler {
	return InventoryHandler{
		InventoryService: inventoryService,
		Config: config,
	}
}

// create inventory
func (inventoryHandler *InventoryHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	var req dto.InventoryCreateRequest
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

	// parsing to model inventory
	inventory := model.Inventory{
		ProductId: req.ProductId,
		ShelveId: req.ShelveId,
		Quantity: req.Quantity,
	}

	// create inventory service
	err = inventoryHandler.InventoryService.AddInventory(&inventory)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created inventory", nil)
}

// get list inventorys
func (inventoryHandler *InventoryHandler) GetListInventorys(w http.ResponseWriter, r *http.Request) {
	// page, err := strconv.Atoi(r.URL.Query().Get("page"))
	// if err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
	// 	return
	// }

	// Get data inventory form service all inventory
	inventorys, err := inventoryHandler.InventoryService.GetListInventorys()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch inventorys: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data inventory", inventorys)

}

// get list inventory by id
func (inventoryHandler *InventoryHandler) GetListInventoryByID(w http.ResponseWriter, r *http.Request) {
	inventoryIDstr := chi.URLParam(r, "inventory_id")

	inventoryID, err := strconv.Atoi(inventoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid inventory ID", nil)
		return
	}

	response, err := inventoryHandler.InventoryService.GetListInventoryById(inventoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "inventory not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data inventory by id", response)
}

// update inventory
func (inventoryHandler *InventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	inventoryIDstr := chi.URLParam(r, "inventory_id")

	inventoryID, err := strconv.Atoi(inventoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param inventory id :"+err.Error(), nil)
		return
	}

	var req dto.InventoryUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := inventoryHandler.InventoryService.GetListInventoryById(inventoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.ProductId != nil {
		existing.ProductId = *req.ProductId

	}

	if req.ShelveId != nil {
		existing.ShelveId = *req.ShelveId
	}

	if req.Quantity != nil {
		existing.Quantity = *req.Quantity
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model inventory
	inventory := model.Inventory{
		ProductId: existing.ProductId,
		ShelveId: existing.ShelveId,
		Quantity: existing.Quantity,
	}

	err = inventoryHandler.InventoryService.UpdateInventory(inventoryID, &inventory)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete inventory
func (inventoryHandler *InventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	inventoryIDstr := chi.URLParam(r, "inventory_id")

	inventoryID, err := strconv.Atoi(inventoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid inventory ID", nil)
		return
	}

	err = inventoryHandler.InventoryService.DeleteInventory(inventoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}