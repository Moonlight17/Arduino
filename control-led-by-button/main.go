/**
 * Created by Moonlight17
 *
 */

package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	but := machine.D7
	but.Configure(machine.PinConfig{Mode: machine.PinInput})

	count := 0

	for {
		led.Set(but.Get())
		time.Sleep(time.Millisecond * 100)
		count++
	}
}
