package main

import (
	"github.com/codegangsta/cli"
	"github.com/geminas/govue/controller"
	"log"
	"os"
)

var (
	Version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "govue"
	app.Usage = "The Vue Components helper"
	app.Author = "Geminas"
	app.Email = "geminas@aliyun.com"
	app.Commands = []cli.Command{
		{
			Name:  "component",
			Usage: "Components actions",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "Create a Vue Component",
					Action: componentCreate,
				},
			},
		},
	}
	app.Run(os.Args)
}

func componentCreate(ctx *cli.Context) {
	arg := ctx.Args()
	if len(arg) != 1 {
		log.Println("Please just tell the name of component")
		return
	}
	controller.CreateComponent(arg[0], "./")
}
