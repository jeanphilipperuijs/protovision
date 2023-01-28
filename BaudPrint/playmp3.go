package baudprint

import (
	"bytes"
	"encoding/binary"
	"os"
	"os/signal"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
)

func Playmp3(fileName string) {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	// create mpg123 decoder instance
	decoder, err := mpg123.NewDecoder("")
	chk_mp3(err)

	chk_mp3(decoder.Open(fileName))
	defer decoder.Close()

	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	chk_mp3(err)
	defer stream.Close()

	chk_mp3(stream.Start())
	defer stream.Stop()
	for {
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		chk_mp3(err)

		chk_mp3(binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out))
		chk_mp3(stream.Write())
		select {
		case <-sig:
			return
		default:
		}
	}
}

func chk_mp3(err error) {
	if err != nil {
		panic(err)

	}
}
