package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/project/config"
	"example.com/project/models"
	"example.com/project/mapper"
	"example.com/project/dto"
)

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User Data"
// @Success 200 {object} dto.UserResponse
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var request dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert DTO → Model
	userModel := mapper.ToUserModel(request)

	config.DB.Create(&userModel)

	// Convert Model → Response DTO
	response := mapper.ToUserResponse(userModel)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags users
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	config.DB.Find(&users)

	// Use mapper instead of manual loop
	response := mapper.ToUserResponseList(users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
