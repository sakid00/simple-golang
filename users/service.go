package users

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{repo}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email

	newUser, err := s.repo.Register(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
