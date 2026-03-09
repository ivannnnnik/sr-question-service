package model

import "time"


type Question struct{
	ID           string    `db:"id"`
    Title        string    `db:"title"`
    Category     string    `db:"category"`
    Difficulty string    `db:"difficulty"`
    CreatedAt    time.Time `db:"created_at"`
}