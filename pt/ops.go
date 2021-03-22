package pt

import (
	"context"
	"fmt"

	"github.com/varunamachi/teak"
	"github.com/varunamachi/teak/pg"
)

const listToTaskQuery = `
	INSERT INTO task_to_list(list_id, task_id) VALUES($1, $2) 
`
const setStatusQuery = `
	UPDATE %s SET status = $1 WHERE id = $1
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

func SetListStatus(gtx context.Context, itemID string, status Status) error {
	query := fmt.Sprintf(setStatusQuery, taskListDataType)
	_, err := pg.Conn().ExecContext(gtx, query, itemID, string(status))
	return teak.LogErrorX("pt.ops",
		"Failed status of task list %s to %v", err, itemID, status)
}

func SetTaskStatus(gtx context.Context, itemID string, status Status) error {
	query := fmt.Sprintf(setStatusQuery, taskDataType)
	_, err := pg.Conn().ExecContext(gtx, query, itemID, string(status))
	return teak.LogErrorX("pt.ops",
		"Failed status of task list %s to %v", err, itemID, status)
}

func UpdateTask(gtx context.Context, task *Task) error {
	return nil
}

func UpdateTaskList(gtx context.Context, tl *TaskList) error {
	return nil
}

func DeleteTask(gtx context.Context, taskID string) error {
	return nil
}

func DeleteTaskList(gtx context.Context, tlID string) error {
	return nil
}
