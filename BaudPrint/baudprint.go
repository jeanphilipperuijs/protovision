package baudprint

import (
	"fmt"
	"math/rand"
	"time"
)

func BaudPrint(message string, baudrate int64, variability int, playmp3 bool, playaiff bool) {

	min := int(baudrate) - variability
	max := int(baudrate) + variability

	for _, v := range message {
		rbr := rand.Intn(max-min) + min
		//fmt.Println("\t", max, min, rbr)
		time.Sleep(time.Duration(time.Second / time.Duration(rbr)))

		if playmp3 {
			Playmp3("./BaudPrint/baud.mp3")
		} else if playaiff {
			Play("./BaudPrint/baud.aiff")
		}

		fmt.Printf(string(v))
	}
	time.Sleep(time.Duration(time.Second / 2))
}
