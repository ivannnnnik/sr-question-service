package service

import (
	"context"

	"github.com/ivannnnnik/sr-question-service/internal/model"
)


type questionRepo interface {
	Create(ctx context.Context, question *model.Question) error
	GetByID(ctx context.Context, id string) (*model.Question, error)
	List(ctx context.Context) ([]model.Question, error)
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

func (svc *QuestionService) GetQuestion(ctx context.Context, id string) (*model.Question, error){
	user, err := svc.repo.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}

	return user, nil

}

func (svc *QuestionService) List(ctx context.Context) ([]model.Question, error){
	questions, err := svc.repo.List(ctx)
	if err != nil{
		return nil, err
	}

	return questions, nil

}