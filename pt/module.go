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

func Initialize(gtx context.Context, app *teak.App) (err error) {
	return nil
}

func Setup(gtx context.Context, app *teak.App) (err error) {
	return nil
}

func Reset(gtx context.Context, app *teak.App) (err error) {
	return nil
}
