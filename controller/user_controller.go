package controller

import (
    "go-api/model"
    "go-api/usecase"
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserController interface {
    Register(c *gin.Context)
    Login(c *gin.Context)
}

type userController struct {
    userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
    return &userController{
        userUsecase: usecase,
    }
}

func (uc *userController) Register(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    registeredUser, err := uc.userUsecase.Register(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    registeredUser.Password = ""
    c.JSON(http.StatusCreated, registeredUser)
}

func (uc *userController) Login(c *gin.Context) {
    var login model.LoginRequest
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := uc.userUsecase.Login(login)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    user.Password = ""
    c.JSON(http.StatusOK, user)
}