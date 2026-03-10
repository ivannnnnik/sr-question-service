package service

import (
	"context"

	"github.com/ivannnnnik/sr-question-service/internal/model"
)


type questionRepo interface {
	Create(ctx context.Context, question *model.Question) error
}

type QuestionService struct{
	repo questionRepo
}

func NewQuestionService(repo questionRepo) *QuestionService{
	return &QuestionService{
		repo: repo,
	}
}


func (svc *QuestionService) Create(ctx context.Context, title, category, difficulty string) (*model.Question, error){
	
	questionModel := model.Question{
		Title: title,
		Category: category,
		Difficulty: difficulty,
	}
	
	err := svc.repo.Create(ctx, &questionModel)

	if err != nil{
		return nil, err
	}

	return &questionModel, nil

}