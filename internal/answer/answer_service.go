package answer

import (
	"course/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AnswerService struct {
	db *gorm.DB
}

func NewAnswerService(database *gorm.DB) *AnswerService {
	return &AnswerService{
		db: database,
	}
}

func (answ AnswerService) CreateAnswer(ctx *gin.Context) {
	var input domain.AnswerRequest
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

	userID := ctx.Request.Context().Value("user_id").(float64)

	answer := domain.Answer{
		ExerciseID: input.ExerciseID,
		QuestionID: input.QuestionID,
		Answer:     input.Answer,
		UserID:     int(userID),
	}

	err = answ.db.Create(&answer).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal server while create user",
		})
		return
	}

	answerResponse := domain.AnswerResponse{
		ExerciseID: answer.ExerciseID,
		QuestionID: answer.QuestionID,
		Answer:     answer.Answer,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": answerResponse,
	})

}
