package monitor_test

import (
	"testing"

	"github.com/Catzkorn/go-blood-glucose/monitor"
	"github.com/shopspring/decimal"
)

func TestMonitorInit(t *testing.T) {

	monitor, err := monitor.New("8", "3.9")

	if err != nil {
		t.Fatalf("monitor errored unexpectedly: %v", err)
	}

	monitor.AddReading("4.6")
	readings := monitor.Readings()

	if len(readings) != 1 {
		t.Fatalf("Did not return number of readings as expected, got: %d, want: %d", len(readings), 1)
	}

	expectedDec, _ := decimal.NewFromString("4.6")
	if !readings[0].Equals(expectedDec) {
		t.Errorf("Reading is does not match inputted reading, got: %s, want: %s", readings[0], expectedDec)
	}

}
