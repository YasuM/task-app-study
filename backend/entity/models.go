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
	CreatedAt time.Time
	Status    int32
}