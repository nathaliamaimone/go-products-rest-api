package usecase

import (
    "go-api/model"
    "go-api/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
    Register(user model.User) (model.User, error)
    Login(login model.LoginRequest) (model.User, error)
}

type userUsecase struct {
    repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
    return &userUsecase{
        repository: repo,
    }
}

func (u *userUsecase) Register(user model.User) (model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return model.User{}, err
    }
    user.Password = string(hashedPassword)
    user.Role = "user"
    return u.repository.CreateUser(user)
}

func (u *userUsecase) Login(login model.LoginRequest) (model.User, error) {
    user, err := u.repository.GetUserByEmail(login.Email)
    if err != nil {
        return model.User{}, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
    if err != nil {
        return model.User{}, err
    }

    return user, nil
}