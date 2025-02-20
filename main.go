package main

import (
	"log"
	"os"

	"weather-CLI/weather"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() > 0 {
						location := cCtx.Args().Get(0)
						weather.GetWeatherReport(location)
				}
				return nil
		},
}

	if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
	}
}