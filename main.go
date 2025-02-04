package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "random-password-generator",
		Usage: "Generate a secure random password",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "length",
				Aliases: []string{"l"},
				Value:   12,
				Usage:   "Length of the password",
			},
			&cli.BoolFlag{
				Name:    "special",
				Aliases: []string{"s"},
				Usage:   "Include special characters",
			},
		},
		Action: func(c *cli.Context) error {
			length := c.Int("length")
			useSpecial := c.Bool("special")

			password, err := generatePassword(length, useSpecial)
			if err != nil {
				return err
			}

			fmt.Println("Generated Password:", password)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
