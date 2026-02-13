package mapper

import (
	"example.com/project/dto"
	"example.com/project/models"
)

func ToUserResponse(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponseList(users []models.User) []dto.UserResponse {
	var response []dto.UserResponse

	for _, user := range users {
		response = append(response, ToUserResponse(user))
	}

	return response
}

func ToUserModel(req dto.CreateUserRequest) models.User {
	return models.User{
		Name:  req.Name,
		Email: req.Email,
	}
}
