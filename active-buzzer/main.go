package main

import (
	"machine"
	"time"
)

func main() {
	//  ACTIVE
	buz := machine.D7
	buz.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//for {
	//	buz.Low()
	//	time.Sleep(time.Millisecond * 500)
	//
	//	buz.Low()
	//	time.Sleep(time.Millisecond * 500)
	//}
	//	Passive
	for {
		//for i := 0; i < 80; i++ {
		//	buz.High()
		//	time.Sleep(time.Millisecond * 4)
		//	buz.Low()
		//	time.Sleep(time.Millisecond * 4)
		//}
		//for i := 0; i < 100; i++ {
		//	buz.High()
		//	time.Sleep(time.Millisecond * 3)
		//	buz.Low()
		//	time.Sleep(time.Millisecond * 3)
		//}
		//for i := 0; i < 110; i++ {
		//	buz.High()
		//	time.Sleep(time.Millisecond * 2)
		//	buz.Low()
		//	time.Sleep(time.Millisecond * 2)
		//}
		//for i := 0; i < 120; i++ {
		//	buz.High()
		//	time.Sleep(time.Millisecond * 1)
		//	buz.Low()
		//	time.Sleep(time.Millisecond * 1)
		//}

		time.Sleep(time.Millisecond * 1000000)
	}
}
