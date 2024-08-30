package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eharris128/tui-sandbox/tui"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sandbox",
		Usage: "A toolkit example that integrates CLI and TUI",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run the example command",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "tui",
						Usage: "Enable TUI mode",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("tui") {
						tui.RunTUI()
					}
					fmt.Println("Running in CLI mode")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
