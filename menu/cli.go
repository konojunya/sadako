package menu

import (
	"github.com/konojunya/sadako/action"
	"github.com/urfave/cli"
)

// Getapp Returning the app that finished setting to main.go
func Getapp() *cli.App {
	app := cli.NewApp()
	config(app)
	app.Commands = getCommands()
	return app
}

func config(app *cli.App) {
	app.Name = "sadako"
	app.Usage = "Sadako will appear when git pull"
	app.Version = "1.0.0"
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "init",
			Usage:  "sadako in your git repository",
			Action: action.Set,
		},
		{
			Name:   "rm",
			Usage:  "remove sadako... :)",
			Action: action.Remove,
		},
	}
}
