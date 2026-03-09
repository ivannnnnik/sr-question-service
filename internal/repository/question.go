package repository

import (
	"context"

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