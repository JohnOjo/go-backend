package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CreateUser godoc
// @Summary Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User Data"
// @Success 200 {object} UserResponse
// @Router /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} UserResponse
// @Router /users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	res, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}