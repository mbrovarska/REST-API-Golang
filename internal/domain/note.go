package models

import "time"

type Note struct {
	ID        int64     
	Title     string     `binding:"required"`
	Text      string     `binding:"required"`
	UserID    int64      `binding:"required"`
	CreatedAt time.Time  `binding:"required"`
	UpdatedAt time.Time
}