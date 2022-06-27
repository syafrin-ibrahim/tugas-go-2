package main

import (
	"course/internal/answer"
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/question"
	"course/internal/user"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.NewDatabaseConnection()
	xrService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)
	qsService := question.NewQuestionService(db)
	ansService := answer.NewAnswerService(db)
	r.GET("/exercise/:id", middleware.Auth(userService), xrService.GetExercise)
	r.POST("/exercise", middleware.Auth(userService), xrService.CreteExercise)
	r.GET("/exercise", middleware.Auth(userService), xrService.GetExerciseAll)
	r.GET("/exercise/:id/score", middleware.Auth(userService), xrService.GetUserScore)
	r.POST("/question", middleware.Auth(userService), qsService.CreateQuestion)
	r.POST("/answer", middleware.Auth(userService), ansService.CreateAnswer)
	r.POST("/register", userService.Register)
	r.POST("/login", userService.Login)
	r.Run(":8000")
}

//contohmicroservices
//https://github.com/ahsanulks/user_service

// func handlerGetExercise(ctx *gin.Context) {

// 	ctx.JSON(200, gin.H{
// 		"pesan": "hello handler",
// 	})
// }
