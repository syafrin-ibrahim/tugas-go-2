package exercise

import (
	"course/internal/domain"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExerciseService struct {
	db *gorm.DB
}

func NewExerciseService(database *gorm.DB) *ExerciseService {
	return &ExerciseService{
		db: database,
	}
}

func (ex ExerciseService) GetExerciseAll(ctx *gin.Context) {
	var exercise []domain.Exercise

	err := ex.db.Find(&exercise).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "not found",
		})
		return
	}
	var responses []domain.ExerciseResponse
	for _, exer := range exercise {
		exerciseResponse := domain.ExerciseResponse{
			Title:       exer.Title,
			Description: exer.Description,
		}

		responses = append(responses, exerciseResponse)

	}

	ctx.JSON(http.StatusOK, gin.H{
		"data exercise ": responses,
	})

}

func (ex ExerciseService) CreteExercise(ctx *gin.Context) {
	var input domain.ExerciseRequest

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

	newExercise := domain.Exercise{
		Title:       input.Title,
		Description: input.Description,
	}

	err = ex.db.Create(&newExercise).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal server while create exercise",
		})
		return
	}

	exerciseResponse := domain.ExerciseResponse{

		Title:       newExercise.Title,
		Description: newExercise.Description,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": exerciseResponse,
	})
}

func (ex ExerciseService) GetExercise(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	var exercise domain.Exercise
	err = ex.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, exercise)

}

func (ex ExerciseService) GetUserScore(ctx *gin.Context) {
	var exercise domain.Exercise
	paramId := ctx.Param("id")
	exerciseID, err := strconv.Atoi(paramId)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}

	err = ex.db.Where("id = ?", exerciseID).Preload("Questions").Take(&exercise).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "record not found",
		})
		return
	}

	userID := ctx.Request.Context().Value("user_id").(float64)
	var answers []domain.Answer

	err = ex.db.Where("exercise_id = ? AND user_id = ?", exerciseID, userID).Find(&answers).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"skor ": "0",
		})
		return
	}

	mapQa := make(map[int]domain.Answer)

	for _, answer := range answers {
		mapQa[answer.QuestionID] = answer
	}

	var score int
	for _, question := range exercise.Questions {
		if strings.EqualFold(question.CorrectAnswer, mapQa[question.ID].Answer) {
			score += question.Score
		}
	}

	ctx.JSON(200, gin.H{
		"skor ": score,
	})

}
