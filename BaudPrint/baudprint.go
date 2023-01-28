package baudprint

import (
	"fmt"
	"math/rand"
	"time"
)

func BaudPrint(message string, baudrate int64) {
	variability := 3
	min := int(baudrate) / variability
	max := int(baudrate) * variability
	for _, v := range message {
		rbr := rand.Intn(max-min) + min
		time.Sleep(time.Duration(time.Second / time.Duration(rbr)))

		//Playmp3("./BaudPrint/baud.mp3")
		//Play("./BaudPrint/baud.aiff")
		fmt.Printf(string(v))
	}
}
