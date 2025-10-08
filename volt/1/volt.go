package volt

import (
	"errors"
	"fmt"
)

var (
	ErrHighVoltage = errors.New("reached maximum safety voltage")
	ErrPowerOutage = errors.New("cannot power on pumps")
)

func PowerOnPV(voltage int) error {
	if voltage == 0 {
		return ErrPowerOutage
	}
	if voltage > 250 {
		return fmt.Errorf("switching the system: %w", ErrHighVoltage)
	}
	return nil
}
