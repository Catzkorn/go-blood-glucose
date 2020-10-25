package monitor

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// Monitor defines a new monitor
type Monitor struct {
	upperLimit decimal.Decimal
	lowerLimit decimal.Decimal
	readings   []decimal.Decimal
}

// New creates a new monitor with upper and lower limits
func New(upperLimit string, lowerLimit string) (Monitor, error) {

	upperLimitDec, err := decimal.NewFromString(upperLimit)
	if err != nil {
		return Monitor{}, fmt.Errorf("failed to parse upper limit: %w", err)
	}

	lowerLimitDec, err := decimal.NewFromString(lowerLimit)
	if err != nil {
		return Monitor{}, fmt.Errorf("failed to parse lower limit: %w", err)
	}

	if upperLimitDec.IsNegative() || lowerLimitDec.IsNegative() {
		return Monitor{}, errors.New("limits cannot be negative")
	}

	if upperLimitDec.IsZero() || lowerLimitDec.IsZero() {
		return Monitor{}, errors.New("limits cannot be 0")
	}
	if upperLimitDec.LessThanOrEqual(lowerLimitDec) {
		return Monitor{}, errors.New("upper limit cannot be less than or equal to lower limit")
	}
	monitor := Monitor{
		upperLimit: upperLimitDec,
		lowerLimit: lowerLimitDec,
	}

	return monitor, nil

}

// AddReading accepts a BG reading
func (m *Monitor) AddReading(reading string) error {

	readingDec, err := decimal.NewFromString(reading)
	if err != nil {
		return fmt.Errorf("failed to parse reading: %w", err)
	}
	if readingDec.IsNegative() {
		return errors.New("reading cannot be negative")
	}

	if readingDec.IsZero() {
		return errors.New("reading cannot be 0")
	}

	m.readings = append(m.readings, readingDec)

	return nil

}

// Readings returns all the readings
func (m *Monitor) Readings() []decimal.Decimal {

	return m.readings
}
