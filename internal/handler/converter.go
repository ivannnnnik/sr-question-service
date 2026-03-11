package handler

import (
	"time"

	questionv1 "github.com/ivannnnnik/sr-proto/gen/go/question/v1"
	"github.com/ivannnnnik/sr-question-service/internal/model"
)

func QuestionToProto(question *model.Question) *questionv1.Question{
	return &questionv1.Question{
		Id: question.ID,
		Title: question.Title,
		Category: question.Category,
		Difficulty: question.Difficulty,
		CreatedAt: question.CreatedAt.Format(time.RFC3339),
	}
}


func QuestionsToProto(questions []model.Question) []*questionv1.Question{
	result := []*questionv1.Question{}
	
	for _, question := range questions{
		result = append(result,
			&questionv1.Question{
				Id: question.ID,
				Title: question.Title,
				Category: question.Category,
				Difficulty: question.Difficulty,
				CreatedAt: question.CreatedAt.Format(time.RFC3339),
		})
	}
	return result
}