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

type SaleHandler struct {
	SaleService service.SaleService
	Config utils.Configuration
}

func NewSaleHandler(saleService service.SaleService, config utils.Configuration) SaleHandler {
	return SaleHandler{
		SaleService: saleService,
		Config: config,
	}
}

// create sale
func (saleHandler *SaleHandler) AddSale(w http.ResponseWriter, r *http.Request) {
	var req dto.SaleCreateRequest
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

	// parsing to model sale
	sale := model.Sale{
		UserId: req.UserId,
		ProductId: req.ProductId,
		Items: req.Items,
		// Price: req.Price,
		// Total: req.Total,
	}

	// create sale service
	err = saleHandler.SaleService.AddSale(&sale)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created sale", nil)
}

// get list sales
func (saleHandler *SaleHandler) GetListSales(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
		return
	}

	// limit pagination
	limit := saleHandler.Config.Limit

	// Get data sale form service all sale
	sales, pagination, err := saleHandler.SaleService.GetListSales(page, limit)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch sales: "+err.Error(), nil)
		return
	}

	utils.ResponsePagination(w, http.StatusOK, "success get data sale", sales, *pagination)

}

// get list sale by id
func (saleHandler *SaleHandler) GetListSaleByID(w http.ResponseWriter, r *http.Request) {
	saleIDstr := chi.URLParam(r, "sale_id")

	saleID, err := strconv.Atoi(saleIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid sale ID", nil)
		return
	}

	response, err := saleHandler.SaleService.GetListSaleById(saleID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "sale not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data sale by id", response)
}

// update sale
func (saleHandler *SaleHandler) UpdateSale(w http.ResponseWriter, r *http.Request) {
	saleIDstr := chi.URLParam(r, "sale_id")

	saleID, err := strconv.Atoi(saleIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param sale id :"+err.Error(), nil)
		return
	}

	var req dto.SaleUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := saleHandler.SaleService.GetListSaleById(saleID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.UserId != nil {
		existing.UserId = *req.UserId
	}

	// if req.ProductId != nil {
	// 	existing.ProductId = *req.ProductId
	// }

	if req.Items != nil {
		existing.Items = *req.Items
	}

	// if req.Price != nil {
	// 	existing.Price = *req.Price
	// }

	// if req.Total != nil {
	// 	existing.Total = *req.Total
	// }

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model sale
	sale := model.Sale{
		UserId: existing.UserId,
		// ProductId: existing.ProductId,
		Items: existing.Items,
		// Price: existing.Price,
		// Total: existing.Total,
	}

	err = saleHandler.SaleService.UpdateSale(saleID, &sale)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete sale
func (saleHandler *SaleHandler) DeleteSale(w http.ResponseWriter, r *http.Request) {
	saleIDstr := chi.URLParam(r, "sale_id")

	saleID, err := strconv.Atoi(saleIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid sale ID", nil)
		return
	}

	err = saleHandler.SaleService.DeleteSale(saleID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}