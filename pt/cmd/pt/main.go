package main

import (
	"context"
	"os"

	"github.com/varunamachi/patti/pt"
	"github.com/varunamachi/teak"
	"github.com/varunamachi/teak/pg"
)

func main() {
	app := teak.NewApp(
		"kmvproxy",
		teak.Version{
			Major: 0,
			Minor: 0,
			Patch: 0,
		},
		0,
		"The KMVProxy ETL & Server",
		teak.DefaultAuthenticator,
		teak.NoOpAuthorizer,
		pg.NewUserStorage(),
		pg.NewStorage())

	// pg.SetDefaultConn("kmvproxy")
	app.AddModule(pt.NewModule())
	app.Commands = append(app.Commands, *teak.GetServiceStartCmd(teak.Serve))
	app.Exec(context.TODO(), os.Args)
}
