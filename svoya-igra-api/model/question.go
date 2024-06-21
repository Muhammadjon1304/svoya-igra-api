package model

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Score    uint   `json:"score"`
}
type Topic struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Question1 Question `json:"question1"`
	Question2 Question `json:"question2"`
	Question3 Question `json:"question3"`
	Question4 Question `json:"question4"`
	Question5 Question `json:"question5"`
}
type PostTopic struct {
	Title     string   `json:"title"`
	Question1 Question `json:"question1"`
	Question2 Question `json:"question2"`
	Question3 Question `json:"question3"`
	Question4 Question `json:"question4"`
	Question5 Question `json:"question5"`
}
type TopicUri struct {
	ID uint `uri:"id" binding:"required,number"`
}
