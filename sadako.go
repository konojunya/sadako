package main

import (
	"os"

	"github.com/konojunya/sadako-git-pull/menu"
)

func main() {
	app := menu.Getapp()
	app.Run(os.Args)
}
