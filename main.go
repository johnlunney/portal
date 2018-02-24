package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"portal/storage"
)

func main() {
	ds := storage.Create()
	app := cli.NewApp()
	app.Name = "portal"
	app.Usage = "With it, you can create your own portals."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "store directory",
			Action: func(c *cli.Context) error {
				ds.Add(c.Args().First())
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list stored directories",
			Action: func(c *cli.Context) error {
				fmt.Println(ds.List())
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		path, err := ds.Match(c.Args().First())
		if err == nil {
			fmt.Print(path)
		} else {
			fmt.Println(err)
		}
		return err
	}

	app.Run(os.Args)
}