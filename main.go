package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var aa = "" +
	"H4sIAGdplloAA3u/Z/97IhBXzeOGBgVcsvHxINkasKIaKwwAFiZWUkH//f7FEJTwqGfS+z170MxW" +
	"gDhlXycQPd0+93Hz3vgiWxu4KRCtYDVAWQgCCe+bBVISHw9UBZIHYoUnO/ZDlb3YNA9IgsTe75sB" +
	"RBDTgD4DmaMPxOjGgRSApLmgcvrggOiHK+ECC2LTzAVUBhTwAglC5CDiQFu5oD6CyCBkQVJcAE4i" +
	"5FepAQAA"

func main() {
	f, err := os.Open("sadako.mp3")
	if err != nil {
		panic(err)
	}

	s, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}

	b, err := base64.StdEncoding.DecodeString(aa)
	if err != nil {
		panic(err)
	}
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, r); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	done := make(chan struct{})
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	_ = <-done
}
