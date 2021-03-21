package pt

import (
	"context"

	"github.com/varunamachi/teak"
	"github.com/varunamachi/teak/pg"
)

const listToTaskQuery = `
	INSERT INTO task_to_list(list_id, task_id) VALUES($1, $2) 
`

// CreateTask - creates a task in the given task list in the database
func CreateTask(gtx context.Context, list string, task *Task) error {
	err := teak.GetStore().Create(gtx, taskDataType, task)
	if err != nil {
		return teak.LogErrorX("pt.ops",
			"Failed to create task %s", err, task.ID)
	}
	_, err = pg.Conn().ExecContext(gtx, listToTaskQuery, list, task.ID)
	return teak.LogErrorX("pt.ops",
		"Failed to add task %s to task list %s", err, task.ID, list)
}

func CreateTaskList(gtx context.Context, tl *TaskList) error {
	return teak.GetStore().Create(gtx, taskListDataType, tl)
}
