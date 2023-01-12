package main

import (
	"fmt"
	"log"

	"github.com/nagygr/backlight/pkg/hw"
)

const (
	brightnessCmd       = "/sys/class/backlight/intel_backlight/brightness"
	actualBrightnessCmd = "/sys/class/backlight/intel_backlight/actual_brightness"
	maxBrightnessCmd    = "/sys/class/backlight/intel_backlight/max_brightness"
)

func main() {
	var (
		brightness    int
		maxBrightness int
		err           error

		brightnessCtrl = hw.NewBrightnessController(
			brightnessCmd, actualBrightnessCmd, maxBrightnessCmd, 0,
		)
	)

	brightness, err = brightnessCtrl.CurrentBrightness()
	if err != nil {
		log.Fatalf("Error acquiring brightness: %s", err)
	}

	fmt.Printf("Brightness: %d\n", brightness)

	maxBrightness, err = brightnessCtrl.MaxBrightness()
	if err != nil {
		log.Fatalf("Error acquiring max brightness: %s", err)
	}

	fmt.Printf("Max brightness: %d\n", maxBrightness)

	err = brightnessCtrl.SetBrightness(1400)
	if err != nil {
		log.Fatalf("Error setting brightness: %s", err)
	}
}
