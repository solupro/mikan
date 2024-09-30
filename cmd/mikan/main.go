package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "mikan",
		Usage: "yet another MiService CLI",
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
