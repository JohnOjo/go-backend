package user

func toUserModel(req CreateUserRequest) User {
	return User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func toUserResponse(user User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func toUserResponseList(users []User) []UserResponse {
	var response []UserResponse

	for _, u := range users {
		response = append(response, toUserResponse(u))
	}

	return response
}
