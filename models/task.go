package models

import (
	"github.com/fatih/structs"
	"time"
)

type Task struct {
	ID            string    `bson:"_id" json:"id" db:"id" structs:"id"`
	Name          string    `bson:"name" json:"name" db:"name" structs:"name"`
	Command       string    `bson:"command" json:"command" db:"command" structs:"command"`
	ExecutionTime time.Time `bson:"execution_time" json:"ExecutionTime" db:"ExecutionTime" structs:"executionTime"`
	CreatedAt     time.Time `bson:"created_at" json:"CreatedAt" db:"CreatedAt" structs:"createdAt"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updatedAt" db:"updatedAt" structs:"updatedAt"`
}

func (c *Task) Map() map[string]interface{} {
	return structs.Map(c)
}

func (c *Task) GetNames() []string {
	fieldNames := structs.Fields(c)
	fieldStructNames := make([]string, len(fieldNames))

	for i, fieldName := range fieldNames {
		fieldString := fieldName.Name()
		fieldStructName := fieldName.Tag(structs.DefaultTagName)
		if fieldStructName == "" {
			fieldStructName = fieldString
		}

		fieldStructNames[i] = fieldStructName
	}

	return fieldStructNames
}
