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

type UserHandler struct {
	UserService service.UserService
	Config utils.Configuration
}

func NewUserHandler(userService service.UserService, config utils.Configuration) UserHandler {
	return UserHandler{
		UserService: userService,
		Config: config,
	}
}

// create user
func (userHandler *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserCreateRequest
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

	// parsing to model user
	user := model.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
		Role: req.Role,
	}

	// create user service
	err = userHandler.UserService.AddUser(&user)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created user", nil)
}

// get list categories
func (userHandler *UserHandler) GetListUsers(w http.ResponseWriter, r *http.Request) {
	// page, err := strconv.Atoi(r.URL.Query().Get("page"))
	// if err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
	// 	return
	// }

	// Get data categori form service all user
	users, err := userHandler.UserService.GetListUsers()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch users: "+err.Error(), nil)
		return
	}

	responseList := utils.ConvertToUserResponseList(users) 

	utils.ResponseSuccess(w, http.StatusOK, "success get data user", responseList)

}

// get list user by id
func (userHandler *UserHandler) GetListUserByID(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	response, err := userHandler.UserService.GetListUserById(userID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "user not found: "+err.Error(), nil)
		return
	}

	responseDTO := utils.ConvertToUserResponse(response)

	utils.ResponseSuccess(w, http.StatusOK, "success get data user by id", responseDTO)
}

// update user
func (userHandler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error param user id :"+err.Error(), nil)
		return
	}

	var req dto.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data :"+err.Error(), nil)
		return
	}

	existing, err := userHandler.UserService.GetListUserById(userID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Data not found:"+err.Error(), nil)
		return
	}

	if req.Username != nil {
		existing.Username = *req.Username
	}

	if req.Email != nil {
		existing.Email = *req.Email
	}

	if req.Password != nil {
		existing.Password = *req.Password
	}

	if req.Role != nil {
		existing.Role = *req.Role
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model user
	user := model.User{
		Username: existing.Username,
		Email: existing.Email,
		Password: existing.Password,
		Role: existing.Role,
	}

	err = userHandler.UserService.UpdateUser(userID, &user)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Error update :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Updated Success", nil)
}

// delete user
func (userHandler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "user_id")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	err = userHandler.UserService.DeleteUser(userID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Error delete :"+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Deleted Success", nil)
}