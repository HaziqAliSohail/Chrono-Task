package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_Map(test *testing.T) {
	type taskFields struct {
		ID            string
		Name          string
		Command       string
		ExecutionTime time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	testCases := []struct {
		name   string
		fields taskFields
		want   map[string]interface{}
	}{
		{
			name: "Student Account struct mapped successfully!",
			fields: taskFields{
				ID:            "3",
				Name:          "Device Data",
				Command:       "lsblk",
				ExecutionTime: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				CreatedAt:     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: map[string]interface{}{
				"id":            "3",
				"name":          "Device Data",
				"command":       "lsblk",
				"executionTime": time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"createdAt":     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				"updatedAt":     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Task{
				ID:            testCase.fields.ID,
				Name:          testCase.fields.Name,
				Command:       testCase.fields.Command,
				ExecutionTime: testCase.fields.ExecutionTime,
				CreatedAt:     testCase.fields.CreatedAt,
				UpdatedAt:     testCase.fields.UpdatedAt,
			}
			if got := testItem.Map(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Map() = %v,\n want %v", got, testCase.want)
			}
		})
	}
}

func TestAccount_GetNames(test *testing.T) {
	type taskFields struct {
		ID            string
		Name          string
		Command       string
		ExecutionTime time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
	testCases := []struct {
		name   string
		fields taskFields
		want   []string
	}{
		{
			name: "Successfully returned the list of account fields",
			fields: taskFields{
				ID:            "3",
				Name:          "Device Data",
				Command:       "lsblk",
				ExecutionTime: time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				CreatedAt:     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
				UpdatedAt:     time.Date(2023, 8, 17, 11, 43, 00, 00, time.UTC),
			},
			want: []string{"id", "name", "command", "executionTime", "createdAt", "updatedAt"},
		},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(subTest *testing.T) {
			testItem := &Task{
				ID:            testCase.fields.ID,
				Name:          testCase.fields.Name,
				Command:       testCase.fields.Command,
				ExecutionTime: testCase.fields.ExecutionTime,
				CreatedAt:     testCase.fields.CreatedAt,
				UpdatedAt:     testCase.fields.UpdatedAt,
			}
			if got := testItem.GetNames(); !reflect.DeepEqual(got, testCase.want) {
				test.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
