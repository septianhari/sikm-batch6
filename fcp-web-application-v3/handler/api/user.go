package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	// TODO: answer here
	var user model.UserLogin

	// Membaca data dari body request
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid decode json"})
		return
	}

	// Memeriksa apakah email atau password kosong
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid decode json"})
		return
	}

	var recordUser = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	// Memanggil service untuk melakukan login user
	userID, err := u.userService.Login(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Membuat token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(recordUser.ID), // UserID sebagai subject
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Membuat atau memperbarui cookie dengan nama session_token
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	// Mengembalikan response dengan status code 200 dan data user yang sudah login
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"message": "login success",
	})
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// TODO: answer here
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika cookie tidak ditemukan, kirim respons status 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Memanggil service untuk mendapatkan daftar tugas pengguna
	tasks, err := u.userService.GetUserTaskCategory()
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data tugas, kirim respons status 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	// Mengembalikan respons dengan daftar tugas pengguna dalam format JSON
	c.JSON(http.StatusOK, tasks)
}
