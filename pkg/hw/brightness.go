package hw

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type BrightnessController struct {
	ActualBrightnessCmd string
}

func NewBrightnessController(actualBrightnessCmd string) *BrightnessController {
	return &BrightnessController{
		ActualBrightnessCmd: actualBrightnessCmd,
	}
}

func (b *BrightnessController) CurrentBrightness() (int, error) {
	var (
		actualBrightnessCommand = exec.Command(
			"cat", b.ActualBrightnessCmd,
		)
		output []byte
		err    error
	)

	output, err = actualBrightnessCommand.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("Couldn't run actual brightness command: %s", err)
	}

	var (
		brightnessStr = strings.TrimSpace(string(output))
		brightness    int
	)
	brightness, err = strconv.Atoi(brightnessStr)
	if err != nil {
		return 0, fmt.Errorf(
			"Couldn't convert the brightness output (%s) to a number: %s",
			brightnessStr, err,
		)
	}

	return brightness, nil
}
