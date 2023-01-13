package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/nagygr/backlight/pkg/hw"
)

func main() {
	var (
		commandDirFlag = flag.String(
			"cmd", "",
			"The name of the directory containing the brightness commands, "+
				"i.e. a directory under \"/sys/class/backlight\" (optional).",
		)
		percentageFlag = flag.Int(
			"p", 0, "The percentage with which the backlight brightness shall be increased/decreased. "+
				"If omitted: the current value is printed.",
		)
		commandRoot string
		err         error
	)

	flag.Parse()

	commandRoot, err = hw.GetBrightnessCommandRoot(*commandDirFlag)
	if err != nil {
		log.Fatalf(
			"Couldn't find brightness commands: %s, please use the -cmd flag",
			err,
		)
	}

	var (
		brightnessCtrl = hw.NewBrightnessController(
			filepath.Join(commandRoot, "brightness"),
			filepath.Join(commandRoot, "actual_brightness"),
			filepath.Join(commandRoot, "max_brightness"),
		)
		currentBrightness int
		maxBrightness     int
	)

	if *percentageFlag == 0 {
		currentBrightness, err = brightnessCtrl.CurrentBrightness()
		if err != nil {
			log.Fatalf("Couldn't acquire current brightness: %s", err)
		}

		maxBrightness, err = brightnessCtrl.MaxBrightness()
		if err != nil {
			log.Fatalf("Couldn't acquire max brightness: %s", err)
		}
	} else {
		currentBrightness, maxBrightness, err = brightnessCtrl.SetPercentage(*percentageFlag)
		if err != nil {
			log.Fatalf("Error setting brightness percentage: %s", err)
		}
	}

	fmt.Printf("Backlight brightness: %d (maximum value: %d)\n", currentBrightness, maxBrightness)
}
