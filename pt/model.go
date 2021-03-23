package pt

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	taskListDataType = "task_lists"
	taskDataType     = "tasks"
)

const (
	defaultListID = "default"
)

type Status string

const (
	Disabled  Status = "Disabled"
	Active    Status = "Active"
	Done      Status = "Done"
	Abandoned Status = "Abandoned"
	OnHold    Status = "OnHold"
)

type Item struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"userID" db:"user_id"`
	Heading     string    `json:"heading" db:"heading"`
	Description string    `json:"description" db:"description"`
	Status      Status    `json:"status" db:"status"`
	CreatedOn   time.Time `json:"createdOn" db:"created_on"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	ModifiedOn  time.Time `json:"modifiedOn" db:"modified_on"`
	ModifiedBy  string    `json:"modifiedBy" db:"modified_by"`
}

// TaskItem - represents a todo item
type Task struct {
	Item
	ExpiresOn time.Time `json:"expiresOn" db:"expires_on"`
}

type TaskList struct {
	Item
	// Tasks []*Task `json:"tasks" db:"tasks"`
}

//TaskItemHandler - CRUD support for Task data type
type TaskItemHandler struct {
}

//DataType - type of data for which this handler is written
func (th *TaskItemHandler) DataType() string {
	return taskDataType

}

//UniqueKeyField - gives the field which uniquely identifies the task
func (th *TaskItemHandler) UniqueKeyField() string {
	return "id"
}

//GetKey - get the uniquely identifying key for the given item
func (th *TaskItemHandler) GetKey(item interface{}) interface{} {
	if agent, ok := item.(Task); ok {
		return agent.ID
	}
	return ""
}

//SetModInfo - set the modifincation information for the data
func (th *TaskItemHandler) SetModInfo(
	item interface{}, at time.Time, by string) {
	if agent, ok := item.(Task); ok {
		agent.ModifiedOn = at
		agent.ModifiedBy = by
	}
}

//CreateInstance - create instance of the data type for which the handler is
//written
func (th *TaskItemHandler) CreateInstance(by string) interface{} {
	return &Task{
		Item: Item{
			ID:         uuid.NewV4().String(),
			CreatedOn:  time.Now(),
			CreatedBy:  by,
			ModifiedOn: time.Now(),
			ModifiedBy: by,
			Status:     Disabled,
		},
	}
}

//PropNames - get prop names of Task struct
func (th *TaskItemHandler) PropNames() []string {
	return []string{
		"id",
		"user_id",
		"heading",
		"description",
		"status",
		"created_on",
		"created_by",
		"modified_on",
		"modified_by",
	}
}

//TaskItemHandler - CRUD support for Task data type
type TaskListHandler struct {
}

//DataType - type of data/table for which this handler is written
func (th *TaskListHandler) DataType() string {
	return taskListDataType

}

//UniqueKeyField - gives the field which uniquely identifies the task list
func (th *TaskListHandler) UniqueKeyField() string {
	return "id"
}

//GetKey - get the uniquely identifying key for the given item
func (th *TaskListHandler) GetKey(item interface{}) interface{} {
	if tl, ok := item.(Task); ok {
		return tl.ID
	}
	return ""
}

//SetModInfo - set the modifincation information for the data
func (th *TaskListHandler) SetModInfo(
	item interface{}, at time.Time, by string) {
	if tlist, ok := item.(Task); ok {
		tlist.ModifiedOn = at
		tlist.ModifiedBy = by
	}
}

//CreateInstance - create instance of the data type for which the handler is
//written
func (th *TaskListHandler) CreateInstance(by string) interface{} {
	return &TaskList{
		Item: Item{
			ID:         uuid.NewV4().String(),
			CreatedOn:  time.Now(),
			CreatedBy:  by,
			ModifiedOn: time.Now(),
			ModifiedBy: by,
			Status:     Disabled,
		},
	}
}

//PropNames - get prop names of TaskList
func (th *TaskListHandler) PropNames() []string {
	return []string{
		"id",
		"user_id",
		"heading",
		"description",
		"status",
		"created_on",
		"created_by",
		"modified_on",
		"modified_by",
	}
}
