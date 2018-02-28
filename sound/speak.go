package sound

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Speak 貞子の音声（git pull~を鳴らす）
func Speak(f *os.File) error {
	s, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	done := make(chan struct{})
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	_ = <-done

	return nil
}
