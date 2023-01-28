package baudprint

import (
	"fmt"
	"math/rand"
	"time"
)

func BaudPrint(message string, baudrate int64, variability int, playmp3 bool, playaiff bool) {

	max_rnd := int(baudrate) / variability
	min_rnd := int(baudrate) * variability

	for _, v := range message {
		rbr := rand.Intn(min_rnd-max_rnd) + max_rnd
		time.Sleep(time.Duration(time.Second / time.Duration(rbr)))

		if playmp3 {
			Playmp3("./BaudPrint/baud.mp3")
		} else if playaiff {
			Play("./BaudPrint/baud.aiff")
		}

		fmt.Printf(string(v))
	}
}
