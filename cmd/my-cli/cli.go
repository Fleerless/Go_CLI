package main

import (
	"os"
	"log"
	"github.com/urfave/cli"
)

func main () {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Lets you query IPs, CNAMES, MX records and Name Servers"

	MyFlags := []cli.Flag{
		cli.StringFlag {
			Name: "Host",
			Value: "jeffreyfleer.com",
		},
	}

	app.Commands = []cli.Command{
		Name: "ns"
		Usage: "Looks up the Name Servers for a Particular Host"
	}
}