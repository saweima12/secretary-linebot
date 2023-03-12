package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/saweima12/secretary-bot/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:        "Secretary-Bot",
		Usage:       "A recording bot for Line message API",
		Description: "",
	}

	app.Commands = []*cli.Command{
		&cmd.CmdServer,
	}

	app.Run(os.Args)
}
