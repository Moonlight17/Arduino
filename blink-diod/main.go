/**
 * Created by Moonlight17
 *
 */

package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 300)

		led.High()
		time.Sleep(time.Millisecond * 700)
		println("teast")
	}
}
