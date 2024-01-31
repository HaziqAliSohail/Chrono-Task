package models

import "time"

type Task struct {
	ID            string    `bson:"_id" json:"id" db:"id"`
	Name          string    `bson:"name" json:"name" db:"name"`
	Command       string    `bson:"command" json:"command" db:"command"`
	ExecutionTime time.Time `bson:"execution_time" json:"ExecutionTime" db:"ExecutionTime"`
	CreatedAt     time.Time `bson:"created_at" json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updatedAt" db:"updatedAt"`
}
