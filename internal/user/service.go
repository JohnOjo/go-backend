package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req CreateUserRequest) (UserResponse, error) {

	user := toUserModel(req)

	if err := s.repo.Create(&user); err != nil {
		return UserResponse{}, err
	}

	return toUserResponse(user), nil
}

func (s *Service) GetUsers() ([]UserResponse, error) {

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return toUserResponseList(users), nil
}
