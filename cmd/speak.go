package main

import (
	"os"

	"github.com/konojunya/sadako-git-pull/art"
	"github.com/konojunya/sadako-git-pull/sound"
)

func main() {
	f, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/konojunya/sadako/media/sadako.mp3")
	if err != nil {
		panic(err)
	}

	art.DrawAA()
	if err = sound.Speak(f); err != nil {
		panic(err)
	}
}
