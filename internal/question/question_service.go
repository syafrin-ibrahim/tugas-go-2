package question

import (
	"course/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(database *gorm.DB) *QuestionService {
	return &QuestionService{
		db: database,
	}
}

func (qs QuestionService) CreateQuestion(ctx *gin.Context) {
	var input domain.QuestionRequest
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

	question := domain.Question{
		ExerciseID:    input.ExerciseID,
		Body:          input.Body,
		OptionA:       input.OptionA,
		OptionB:       input.OptionB,
		OptionC:       input.OptionC,
		OptionD:       input.OptionD,
		CorrectAnswer: input.CorrectAnswer,
		Score:         input.Score,
		CreatorID:     int(userID),
	}

	err = qs.db.Create(&question).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal server while create user",
		})
		return
	}

	questionResponse := domain.QuestionResonse{
		ExerciseID:    question.ExerciseID,
		Body:          question.Body,
		OptionA:       question.OptionA,
		OptionB:       question.OptionB,
		OptionC:       question.OptionC,
		OptionD:       question.OptionD,
		CorrectAnswer: question.CorrectAnswer,
		Score:         question.Score,
		CreatorID:     int(userID),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": questionResponse,
	})
}
