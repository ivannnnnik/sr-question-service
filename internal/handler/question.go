package handler

import (
	"context"

	questionv1 "github.com/ivannnnnik/sr-proto/gen/go/question/v1"
	"github.com/ivannnnnik/sr-question-service/internal/model"
)

type questionService interface {
    Create(ctx context.Context, title, category, difficulty string) (*model.Question, error)
	GetQuestion(ctx context.Context, id string) (*model.Question, error)
	List(ctx context.Context) ([]model.Question, error)
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

func (h *QuestionHandler) GetQuestion(ctx context.Context, req *questionv1.GetQuestionRequest) (*questionv1.GetQuestionResponse, error) {
	question, err := h.service.GetQuestion(ctx, req.Id)
	if err != nil{
		return nil, err
	}

	questionConv := QuestionToProto(question)

	return &questionv1.GetQuestionResponse{
		Question: questionConv,
	}, nil

}

func (h *QuestionHandler) ListQuestions(ctx context.Context, req *questionv1.ListQuestionsRequest) (*questionv1.ListQuestionsResponse, error) {
	questions, err := h.service.List(ctx)
	if err != nil{
		return nil, err
	}

	questionsConv := QuestionsToProto(questions)

	return &questionv1.ListQuestionsResponse{
		Questions: questionsConv,
	}, nil

}