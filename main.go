package main

import (
	"os"

	"csvegeta/cmd"
	"github.com/urfave/cli"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "csvegeta"

	app.Commands = []cli.Command{
		{
			Name:   "headings",
			Usage:  "output the column headings",
			Action: cmd.Headings,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
