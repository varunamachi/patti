package pt

import "time"

type Status string

const (
	Active    Status = "Active"
	Done      Status = "Done"
	Abandoned Status = "Abandoned"
	OnHold    Status = "OnHold"
)

type Item struct {
	ID          int       `json:"id" db:"id"`
	Heading     string    `json:"heading" db:"heading"`
	Description string    `json:"description" db:"description"`
	Status      Status    `json:"status" db:"status"`
	Created     time.Time `json:"created" db:"created"`
	Modified    time.Time `json:"modified" db:"modified"`
}

// TaskItem - represents a todo item
type Task struct {
	Item
	Deadline time.Time `json:"deadline" db:"deadline"`
}

type TaskList struct {
	Item
	Tasks []*Task `json:"tasks" db:"tasks"`
}
