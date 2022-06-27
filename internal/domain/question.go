package domain

import "time"

type Question struct {
	ID            int
	ExerciseID    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	CreatorID     int
	createdAt     time.Time
	updatedAt     time.Time
}

type QuestionRequest struct {
	ExerciseID    int    `json:"exercise_id" binding:"required"`
	Body          string `json:"body" binding:"required"`
	OptionA       string `json:"option_a" binding:"required"`
	OptionB       string `json:"option_b" binding:"required"`
	OptionC       string `json:"option_c" binding:"required"`
	OptionD       string `json:"option_d" binding:"required"`
	CorrectAnswer string `json:"correct_answer" binding:"required"`
	Score         int    `json:"score" binding:"required"`
	createdAt     time.Time
	updatedAt     time.Time
}

type QuestionResonse struct {
	ID            int    `json:"id_question"`
	ExerciseID    int    `json:"exercise id"`
	Body          string `json:"body"`
	OptionA       string `json:"option a"`
	OptionB       string `json:"option b"`
	OptionC       string `json:"option c"`
	OptionD       string `json:"option d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
	CreatorID     int    `json:"creator_id"`
}
