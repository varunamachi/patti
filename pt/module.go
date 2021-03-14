package pt

import (
	"context"

	"github.com/varunamachi/teak"
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
	return nil
}

func Setup(gtx context.Context, app *teak.App) (err error) {
	return nil
}

func Reset(gtx context.Context, app *teak.App) (err error) {
	return nil
}
