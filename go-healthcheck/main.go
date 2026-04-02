package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Health checker",
		Usage: "A tiny tool that checks a website is running or down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			port := cmd.String("port")
			if cmd.String("port") == "" {
				port = "80"
			}
			status := Check(cmd.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// go run . --d discord.com
