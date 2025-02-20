package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"localhost.com/weather-API/weather"
)

func main() {
	var lang string

	app := &cli.App{
		Flags: []cli.Flag{
				&cli.StringFlag{
						Name:  "lang",
						Aliases: []string{"l", "L"},
						Value: "english",
						Usage: "language for the greeting",
						Destination: &lang,
				},
		},
		Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() > 0 {
						location := cCtx.Args().Get(0)
						weather.GetWeatherReport(location)
				}
				if lang != "english" {
						fmt.Println("Lang set to", lang)
				}
				return nil
		},
}

	if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
	}
}