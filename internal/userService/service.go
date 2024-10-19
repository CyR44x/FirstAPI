package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) PostUser(user User) (User, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PatchUserByID(id int, user User) (User, error) {
	return s.repo.PatchUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}
