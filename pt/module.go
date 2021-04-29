package pt

import (
	"context"
	"fmt"
	"time"

	"github.com/varunamachi/teak"
	"github.com/varunamachi/teak/pg"
)

func NewModule() *teak.Module {

	// TODO: - add details
	return &teak.Module{
		Name:        "Patti",
		Description: "Todo list server",
		Endpoints:   []*teak.Endpoint{},
		Commands:    getCommands(),
		Initialize:  Initialize,
		Setup:       Setup,
		Reset:       Reset,
		ItemHandlers: []teak.StoredItemHandler{
			&TaskItemHandler{},
			&TaskListHandler{},
		},
	}
}

var tables = []struct {
	name  string
	query string
}{
	{
		name: "tasks",
		query: `
		CREATE TABLE task(
			id CHAR(32) PRIMARY KEY,
			user_id VARCHAR(256),
			heading VARCHAR(256),
			description TEXT,
			status CHAR(64),
			created TIMESTAMPTZ,
			deadline TIMESTAMPTZ,
			modified TIMESTAMPTZ
		);
		`,
	},
	{
		name: "tasklist",
		query: `
		CREATE TABLE tasklist(
			id CHAR(32) PRIMARY KEY,
			user_id VARCHAR(256),
			heading VARCHAR(256),
			description TEXT,
			status CHAR(64),
			created TIMESTAMPTZ,
			modified TIMESTAMPTZ
		);
		`,
	},
	{
		name: "task_to_list",
		query: `
		CREATE TABLE task_to_list(
			task_id CHAR(32),
			list_id CHAR(32),
			FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE,
			FOREIGN KEY (list_id) REFERENCES tasklist(id) ON DELETE CASCADE,
			PRIMARY KEY(task_id, list_id)
		);
		`,
	},
}

func Initialize(gtx context.Context, app *teak.App) (err error) {
	for _, tab := range tables {
		_, err = pg.Conn().ExecContext(gtx, tab.query)
		if err != nil {
			err = teak.LogErrorX("pt.pg.store", "Failed to create table '%s'",
				err, tab.name)
			break
		}
	}
	if err != nil {
		return err
	}
	CreateTaskList(gtx, &TaskList{
		Item: Item{
			// ID:          uuid.NewV4().String(),
			Heading:     "Default",
			Description: "Default task list",
			Status:      Active,
			CreatedOn:   time.Now(),
			CreatedBy:   "",
			ModifiedOn:  time.Now(),
			ModifiedBy:  "",
		},
		// Tasks: []*Task{},
	})
	return err
}

func Setup(gtx context.Context, app *teak.App) (err error) {
	return err
}

func Reset(gtx context.Context, app *teak.App) (err error) {
	for _, tab := range tables {
		query := fmt.Sprintf("DELETE FROM %s;", tab.name)
		_, err = pg.Conn().ExecContext(gtx, query)
		if err != nil {
			teak.Error(
				"t.pg.store", "Failed clear data from %s: %v", tab.name, err)
			//break??
		}
	}
	return err
}
