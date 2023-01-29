package baudprint

import (
	"fmt"
	"math/rand"
	"time"
)

func BaudPrint(message string, baudrate int64, variability int, playmp3 bool, playaiff bool) {
	bps := baudrate / 10
	min := int(bps) - (variability / 10)
	max := int(bps) + (variability / 10)

	for _, v := range message {
		rbr := rand.Intn(max-min) + min
		//fmt.Println("\t", max, min, rbr)
		time.Sleep(time.Duration(time.Second / time.Duration(rbr)))
		/*
			not working...or maybe it does somewhere?
				if playmp3 {
					Playmp3("./BaudPrint/baud.mp3")
				} else if playaiff {
					Play("./BaudPrint/baud.aiff")
				}
		*/
		fmt.Printf("%c", v)
	}
	//time.Sleep(time.Duration(time.Second / 2))
	time.Sleep(time.Duration(time.Duration(1 / bps)))

}
