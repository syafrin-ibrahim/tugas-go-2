package domain

import "time"

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AnswerRequest struct {
	ExerciseID int    `json:"exercise_id" binding:"required"`
	QuestionID int    `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
}

type AnswerResponse struct {
	ExerciseID int    `json:"exercise_id"`
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}
