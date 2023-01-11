package main

import (
	"fmt"
	"log"

	"github.com/nagygr/backlight/pkg/hw"
)

const (
	actualBrightnessCmd = "/sys/class/backlight/intel_backlight/actual_brightness"
)

func main() {
	var (
		brightnessCtrl = hw.NewBrightnessController(actualBrightnessCmd)
		brightness     int
		err            error
	)

	brightness, err = brightnessCtrl.CurrentBrightness()
	if err != nil {
		log.Fatalf("Error acquiring brightness: %s", err)
	}

	fmt.Printf("Brightness: %d\n", brightness)

}
