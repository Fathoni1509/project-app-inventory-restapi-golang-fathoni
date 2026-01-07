package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"

	"github.com/go-chi/chi/v5"
	"strconv"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
	Config utils.Configuration
}

func NewCategoryHandler(categoryService service.CategoryService, config utils.Configuration) CategoryHandler {
	return CategoryHandler{
		CategoryService: categoryService,
		Config: config,
	}
}

// create category
func (categoryHandler *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CategoryCreateRequest
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

	// parsing to model category
	category := model.Category{
		Name: req.Name,
		Description: req.Description,
	}

	// create category service
	err = categoryHandler.CategoryService.AddCategory(&category)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created category", nil)
}

// get list categories
func (categoryHandler *CategoryHandler) GetListCategories(w http.ResponseWriter, r *http.Request) {
	// Get data categori form service all category
	categories, err := categoryHandler.CategoryService.GetListCategories()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch categories: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data category", categories)

}

// get list category by id
func (categoryHandler *CategoryHandler) GetListCategoryByID(w http.ResponseWriter, r *http.Request) {
	categoryIDstr := chi.URLParam(r, "category_id")

	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	response, err := categoryHandler.CategoryService.GetListCategoryById(categoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Category not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data category by id", response)
}

// update category
func (categoryHandler *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDstr := chi.URLParam(r, "category_id")

	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param category id :"+err.Error(), nil)
		return
	}

	var req dto.CategoryUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := categoryHandler.CategoryService.GetListCategoryById(categoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.Description != nil {
		existing.Description = *req.Description
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model category
	category := model.Category{
		Name: existing.Name,
		Description: existing.Description,
	}

	err = categoryHandler.CategoryService.UpdateCategory(categoryID, &category)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete category
func (categoryHandler *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDstr := chi.URLParam(r, "category_id")

	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	err = categoryHandler.CategoryService.DeleteCategory(categoryID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}