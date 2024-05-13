package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var loginReq model.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	user, err := u.userService.Login(&model.User{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid email or password"))
		return
	}

	token := u.userService.GenerateToken(user.ID)

	c.JSON(http.StatusOK, model.NewSuccessResponse("login success", map[string]string{
		"access_token": token,
	}))
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// Get the access token from the header
	tokenString := c.Request.Header.Get("Authorization")

	if len(tokenString) == 0 {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("access token is missing"))
		return
	}

	// Validate the access token
	claims, err := u.userService.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid access token"))
		return
	}

	// Get the user task category
	userTaskCategory, err := u.userService.GetUserTaskCategory(claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("get user task category success", userTaskCategory))
}
