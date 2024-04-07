package users

import "errors"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginIniput) (User, error)
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

func (s *service) LoginUser(input LoginIniput) (User, error) {
	userLogin, err := s.repo.FindEmail(input.Email)
	if err != nil {
		return userLogin, err
	}

	if userLogin.Id == 0 {
		return userLogin, errors.New("user not found")
	}

	return userLogin, nil
}
