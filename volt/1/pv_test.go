package volt_test

import (
	"errors"
	"testing"

	"github.com/qba73/volt"
)

func TestPowerOnPV_ReturnsPowerOutageErrorForLowVoltage(t *testing.T) {

	err := volt.PowerOnPV(0)
	want := volt.ErrPowerOutage

	if !errors.Is(err, want) {
		t.Errorf("want %, got %T", want, err)
	}
}
