package user

import (
	"course/internal/domain"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		db: database,
	}
}

func (us UserService) Register(ctx *gin.Context) {
	var input domain.RegisterRequest

	err := ctx.ShouldBind(&input)

	if err != nil {

		errorMesssages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			//message := fmt.Sprintf("Error on field %s condition : %s ", e.Field(), e.ActualTag())
			message := fmt.Sprintf("%s tidak boleh kosong ", e.Field())
			errorMesssages = append(errorMesssages, message)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errorMesssages,
		})
		return

	}

	if len(input.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "password minimum 6 character",
		})

		return

	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		NoHp:     input.Phone,
	}

	err = us.db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal server while create user",
		})
		return
	}
	token, err := GenerateToken(user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal server while create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (us UserService) Login(ctx *gin.Context) {
	var user domain.User
	var input domain.LoginRequest

	err := ctx.ShouldBind(&input)

	if err != nil {

		errorMesssages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			//message := fmt.Sprintf("Error on field %s condition : %s ", e.Field(), e.ActualTag())
			message := fmt.Sprintf("%s tidak boleh kosong ", e.Field())
			errorMesssages = append(errorMesssages, message)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errorMesssages,
		})
		return

	}

	err = us.db.Where("email = ?", input.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email tiak terdaftar dalam sistem",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "password salah.....",
		})
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error generate token",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token ": token,
	})

}

var signitureKey = []byte("syafrin12")

func GenerateToken(iduser int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": iduser,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "syafrin",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signitureKey)
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

func (us UserService) DecriptJwt(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token invalid 1")
		}
		return signitureKey, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("token invalid 2")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil

}
