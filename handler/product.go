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

type ProductHandler struct {
	ProductService service.ProductService
	Config utils.Configuration
}

func NewProductHandler(productService service.ProductService, config utils.Configuration) ProductHandler {
	return ProductHandler{
		ProductService: productService,
		Config: config,
	}
}

// create product
func (productHandler *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.ProductCreateRequest
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

	// parsing to model product
	product := model.Product{
		Name: req.Name,
		CategoryId: req.CategoryId,
		ShelveId: req.ShelveId,
		PurchasePrice: req.PurchasePrice,
		SellPrice: req.SellPrice,
		Quantity: req.Quantity,
		UpdatedBy: req.UpdatedBy,
	}

	// create product service
	err = productHandler.ProductService.AddProduct(&product)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created product", nil)
}

// get list products
func (productHandler *ProductHandler) GetListProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
		return
	}

	// limit pagination
	limit := productHandler.Config.Limit

	// Get data product form service all product
	products, pagination, err := productHandler.ProductService.GetListProducts(page, limit)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch products: "+err.Error(), nil)
		return
	}

	utils.ResponsePagination(w, http.StatusOK, "success get data product", products, *pagination)

}

// get list product by id
func (productHandler *ProductHandler) GetListProductByID(w http.ResponseWriter, r *http.Request) {
	productIDstr := chi.URLParam(r, "product_id")

	productID, err := strconv.Atoi(productIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	response, err := productHandler.ProductService.GetListProductById(productID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "product not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data product by id", response)
}

// update product
func (productHandler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productIDstr := chi.URLParam(r, "product_id")

	productID, err := strconv.Atoi(productIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param product id :"+err.Error(), nil)
		return
	}

	var req dto.ProductUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := productHandler.ProductService.GetListProductById(productID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.CategoryId != nil {
		existing.CategoryId = *req.CategoryId
	}

	if req.ShelveId != nil {
		existing.ShelveId = *req.ShelveId
	}

	if req.PurchasePrice != nil {
		existing.PurchasePrice = *req.PurchasePrice
	}

	if req.SellPrice != nil {
		existing.SellPrice = *req.SellPrice
	}

	if req.Quantity != nil {
		existing.Quantity = *req.Quantity
	}

	if req.UpdatedBy != nil {
		existing.UpdatedBy = *req.UpdatedBy
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model product
	product := model.Product{
		Name: existing.Name,
		CategoryId: existing.CategoryId,
		ShelveId: existing.ShelveId,
		PurchasePrice: existing.PurchasePrice,
		SellPrice: existing.SellPrice,
		Quantity: existing.Quantity,
		UpdatedBy: existing.UpdatedBy,
	}

	err = productHandler.ProductService.UpdateProduct(productID, &product)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete product
func (productHandler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIDstr := chi.URLParam(r, "product_id")

	productID, err := strconv.Atoi(productIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	err = productHandler.ProductService.DeleteProduct(productID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}