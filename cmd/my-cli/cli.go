package main

import (
	"fmt"
	"os"
	"log"
	"net"
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

	app.Commands = []*cli.Command{
		{
		Name: "ns",
		Usage: "Looks up the Name Servers for a Particular Host",
		Flags: MyFlags,
		Action: func(c *cli.Context) error {
			ns, err := net.LookupNS(c.String("host"))
			if err != nil {
				return err 
			}
			for i:=0; i<len(ns); i++ {
				fmt.Println(ns[i].Host)
			}
			return nil
		},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}