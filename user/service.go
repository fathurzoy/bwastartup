package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvaiable(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	userEmail, err := s.repository.FindByEmail(input.Email)
	if err != nil{
		return userEmail, err
	}

	if userEmail.ID != 0{
		return userEmail, errors.New("Email has been registered")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// mapping struct input ke struct user
// simpan struct User melalui respository

func (s *service) Login(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil{
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvaiable(input CheckEmailInput) (bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return false, err
	}

	if user.ID == 0{
		return true, nil
	}

	return false, nil
}
