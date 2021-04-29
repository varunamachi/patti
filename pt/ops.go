package pt

import (
	"context"
	"fmt"
	"time"

	"github.com/varunamachi/teak"
	"github.com/varunamachi/teak/pg"
)

const listToTaskQuery = `
	INSERT INTO task_to_list(list_id, task_id) VALUES($1, $2) 
`
const setStatusQuery = `
	UPDATE %s SET status = $3 WHERE id = $1 AND user_id =$2
`

const getAllTasksQuery = `
	SELECT * FROM tasks WHERE task_id IN (
		SELECT task_id FROM task_to_list WHERE list_id = $1
	) ORDER BY 
`

const getAllTaskListQuery = `
	SELECT * FROM tasklists WHERE user_id = $1
`

const getActiveTaskListsQuery = `
	SELECT * FROM tasklist WHERE user_id = $1 AND status = $2
`

// CreateTask - creates a task in the given task list in the database
func CreateTask(gtx context.Context, listID string, task *Task) error {
	if listID == "" {
		//TODO
		// - Check if default list exists for the user
		// - Create it if it does not exist
		listID = defaultListID
	}

	task.CreatedOn = time.Now()
	task.CreatedBy = task.UserID
	task.ModifiedOn = time.Now()
	task.ModifiedBy = task.UserID

	err := teak.GetStore().Create(gtx, taskDataType, task)
	if err != nil {
		return teak.LogErrorX("pt.ops",
			"Failed to create task %s", err, task.ID)
	}
	_, err = pg.Conn().ExecContext(gtx, listToTaskQuery, listID, task.ID)
	return teak.LogErrorX("pt.ops",
		"Failed to add task %s to task list %s", err, task.ID, listID)
}

func CreateTaskList(gtx context.Context, tl *TaskList) error {
	tl.CreatedOn = time.Now()
	tl.CreatedBy = tl.UserID
	tl.ModifiedOn = time.Now()
	tl.ModifiedBy = tl.UserID
	return teak.GetStore().Create(gtx, taskListDataType, tl)
}

func SetListStatus(
	gtx context.Context, itemID, userID string, status Status) error {
	query := fmt.Sprintf(setStatusQuery, taskListDataType)
	_, err := pg.Conn().ExecContext(gtx, query, itemID, userID, string(status))
	return teak.LogErrorX("pt.ops",
		"Failed status of task list %s to %v", err, itemID, status)
}

func SetTaskStatus(
	gtx context.Context, itemID, userID string, status Status) error {
	query := fmt.Sprintf(setStatusQuery, taskDataType)
	_, err := pg.Conn().ExecContext(gtx, query, itemID, userID, string(status))
	return teak.LogErrorX("pt.ops",
		"Failed status of task list %s to %v", err, itemID, status)
}

// func UpdateTask(gtx context.Context, task *Task) error {
// 	return teak.GetStore().Update(gtx, taskDataType, "id", task.ID, task)
// }

// func UpdateTaskList(gtx context.Context, tl *TaskList) error {
// 	return teak.GetStore().Update(gtx, taskListDataType, "id", tl.ID, tl)
// }

func DeleteTask(gtx context.Context, userID, taskID string) error {
	return teak.GetStore().Delete(gtx, taskDataType, "id", taskID)
}

func DeleteTaskList(gtx context.Context, userID, tlID string) error {
	return teak.GetStore().Delete(gtx, taskListDataType, "id", tlID)
}

func GetAllTaskLists(gtx context.Context, userID string) ([]*TaskList, error) {
	// NEXT
	return nil, nil
}

func GetActiveTaskLists(
	gtx context.Context, userID string) ([]*TaskList, error) {
	// NEXT
	return nil, nil
}

func GetTasks(
	gtx context.Context, userID, taskListID string) ([]*TaskList, error) {
	// NEXT
	return nil, nil
}
