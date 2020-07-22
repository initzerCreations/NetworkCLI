package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let you query IPs, CNAMEs, MX records and Name Servers"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the name servers for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
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
