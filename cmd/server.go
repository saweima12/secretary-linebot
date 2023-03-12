package cmd

import (
	"fmt"

	"github.com/saweima12/secretary-bot/internal/server"
	"github.com/urfave/cli/v2"
)

var CmdServer = cli.Command{
	Name:   "server",
	Action: runServer,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Value:   "8001",
			Usage:   "Temporary port number",
			Aliases: []string{"p"},
		},
	},
}

func runServer(ctx *cli.Context) error {

	app := server.New()

	// Get port number from command flag.
	port := ctx.String("port")
	portStr := fmt.Sprintf(":%v", port)

	err := app.Run(portStr)
	if err != nil {
		return err
	}

	return nil
}
