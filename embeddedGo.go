package main
//done with tinyGo
import (
	"time"
	"machine"
)

func main() {
	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		delay(1000)
	}
}

func delay(t uint32) {
	time.Sleep(time.Duration(1000 * t))
}
