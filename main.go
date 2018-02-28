package main

import (
	"os"

	"github.com/konojunya/sadako/menu"
)

func main() {
	app := menu.Getapp()
	app.Run(os.Args)
}
