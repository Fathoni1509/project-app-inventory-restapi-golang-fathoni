package utils

import (
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
)

func ConvertToUserResponse(u model.User) dto.UserResponse {
    return dto.UserResponse{
		UserId: u.UserId,
		Username: u.Username,
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
    }
}

func ConvertToUserResponseList(users []model.User) []dto.UserResponse {
    var response []dto.UserResponse
    for _, u := range users {
        response = append(response, ConvertToUserResponse(u))
    }
    return response
}