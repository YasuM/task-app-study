// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package entity

import (
	"time"
)

type Task struct {
	ID        int64
	Name      string
	UserID    int64
	CreatedAt time.Time
	Status    int32
}

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
