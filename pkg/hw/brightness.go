package hw

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/cznic/mathutil"
)

type BrightnessController struct {
	brightnessCmd       string
	actualBrightnessCmd string
	maxBrightnessCmd    string
	maxBrightness       int
}

func NewBrightnessController(
	brightnessCmd string,
	actualBrightnessCmd string,
	maxBrightnessCmd string,
) *BrightnessController {
	return &BrightnessController{
		brightnessCmd:       brightnessCmd,
		actualBrightnessCmd: actualBrightnessCmd,
		maxBrightnessCmd:    maxBrightnessCmd,
	}
}

func (b *BrightnessController) CurrentBrightness() (int, error) {
	var (
		actualBrightnessCommand = exec.Command(
			"cat", b.actualBrightnessCmd,
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

func (b *BrightnessController) MaxBrightness() (int, error) {
	if b.maxBrightness == 0 {
		var (
			maxBrightnessCommand = exec.Command(
				"cat", b.maxBrightnessCmd,
			)
			output []byte
			err    error
		)

		output, err = maxBrightnessCommand.CombinedOutput()
		if err != nil {
			return 0, fmt.Errorf("Couldn't run max brightness command: %s", err)
		}

		var (
			maxBrightnessStr = strings.TrimSpace(string(output))
			maxBrightness    int
		)
		maxBrightness, err = strconv.Atoi(maxBrightnessStr)
		if err != nil {
			return 0, fmt.Errorf(
				"Couldn't convert the max brightness output (%s) to a number: %s",
				maxBrightnessStr, err,
			)
		}

		b.maxBrightness = maxBrightness
	}

	return b.maxBrightness, nil
}

func (b *BrightnessController) SetBrightness(value int) error {
	var (
		maxBrightness int
		err           error
	)

	maxBrightness, err = b.MaxBrightness()
	if err != nil {
		return fmt.Errorf("Error acquiring max brightness: %w", err)
	}

	if value < 0 || value > maxBrightness {
		return fmt.Errorf("Value (%d) out of range [0, %d]", value, maxBrightness)
	}

	err = os.WriteFile(b.brightnessCmd, []byte(fmt.Sprintf("%d", value)), 0777)
	if err != nil {
		return fmt.Errorf("Couldn't run brightness command: %s", err)
	}

	return nil
}

func (b *BrightnessController) SetPercentage(value int) (currentValue, maxValue int, err error) {
	if value < -100 || value > 100 {
		err = fmt.Errorf("Value (%d) out of range [-100, 100]", value)
		return
	}

	maxValue, err = b.MaxBrightness()
	if err != nil {
		err = fmt.Errorf("Error acquiring max brightness: %w", err)
		return
	}

	currentValue, err = b.CurrentBrightness()
	if err != nil {
		err = fmt.Errorf("Error acquiring current brightness: %w", err)
		return
	}

	var (
		onePercent  float32 = float32(maxValue) / 100.0
		targetValue int     = currentValue + int(onePercent*float32(value))
	)

	currentValue = mathutil.Clamp(targetValue, 0, maxValue)
	err = b.SetBrightness(currentValue)

	return
}
