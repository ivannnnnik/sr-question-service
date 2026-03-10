package handler

import (
	"context"

	questionv1 "github.com/ivannnnnik/sr-proto/gen/go/question/v1"
	"github.com/ivannnnnik/sr-question-service/internal/model"
)

type questionService interface {
    Create(ctx context.Context, title, category, difficulty string) (*model.Question, error)
}

type QuestionHandler struct{
	questionv1.UnimplementedQuestionServiceServer
	service questionService
}

func NewQuestionHandler(svc questionService) *QuestionHandler{
	return &QuestionHandler{
		service: svc,
	}
}

func (h *QuestionHandler) CreateQuestion(ctx context.Context, req *questionv1.CreateQuestionRequest) (*questionv1.CreateQuestionResponse, error) {
	question, err := h.service.Create(ctx, req.Title, req.Category, req.Difficulty)
	if err != nil{
		return nil, err
	}

	questionConv := QuestionToProto(question)


	return &questionv1.CreateQuestionResponse{
		Question: questionConv,
	}, nil

}