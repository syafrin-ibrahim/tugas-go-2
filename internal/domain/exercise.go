package domain

type Exercise struct {
	ID          int    `json:"id exercise"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Questions   []Question
}

type ExerciseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ExerciseResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
