package repository

import (
	"context"
	"fmt"

	"github.com/ivannnnnik/sr-question-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type QuestionRepository struct{
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) *QuestionRepository{
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) Create(ctx context.Context, question *model.Question) error{

	query := `
	INSERT INTO question(title, category, difficulty)
	VALUES ($1, $2, $3) 
	RETURNING id, title, category, difficulty, created_at;
	`

	err := r.db.QueryRowContext(ctx, query, question.Title, question.Category, question.Difficulty).
	Scan(&question.ID, &question.Title, &question.Category, &question.Difficulty, &question.CreatedAt)
	
	return err


}

func (r *QuestionRepository) GetByID(ctx context.Context, id string) (*model.Question, error){
	query := `SELECT id, title, category, difficulty, created_at FROM question WHERE id = $1`

	var question model.Question
	err := r.db.GetContext(ctx, &question, query, id)
	if err != nil{
		return nil, err
	}

	return &question, nil
}

func (r *QuestionRepository) List(ctx context.Context) ([]model.Question, error){
	query := `SELECT id, title, category, difficulty, created_at FROM question`

 	args := map[string]interface{}{}

	var questions []model.Question
	rows, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil{
        return nil, fmt.Errorf("listing questions: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
		var q model.Question

		if err := rows.StructScan(&q);err != nil{
			return nil, fmt.Errorf("scanning question: %v", err)
		}
		questions = append(questions, q)
	}

	return questions, rows.Err()
}