package controller

import (
    "go-api/model"
    "go-api/usecase"
    "go-api/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type UserController interface {
    Register(c *gin.Context)
    Login(c *gin.Context)
}

type userController struct {
    userUsecase usecase.UserUsecase
    jwtService  service.JWTService
}

func NewUserController(usecase usecase.UserUsecase, jwtService service.JWTService) UserController {
    return &userController{
        userUsecase: usecase,
        jwtService:  jwtService,
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

    token, err := uc.jwtService.GenerateToken(user.Id, user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id": user.Id,
            "email": user.Email,
            "role": user.Role,
        },
    })
}