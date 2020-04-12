package clock

import (
	"fmt"
)

// Type clock retains information of an int.
type Clock int

const MINUTES_IN_A_DAY = 1440

func (c Clock) Verify() Clock {
	for c >= MINUTES_IN_A_DAY {
		c -= MINUTES_IN_A_DAY
	}

	for c < 0 {
		c += MINUTES_IN_A_DAY
	}

	return c
}

// New creates a new clock with given hours and minutes.
func New(hours, minutes int) Clock {
	return Clock(hours*60 + minutes).Verify()
}

// Add a given number of minutes to clock.
func (c Clock) Add(minutes int) Clock {
	return Clock(int(c) + minutes).Verify()
}

// Add a given number of minutes to clock.
func (c Clock) Subtract(minutes int) Clock {
	return Clock(int(c) - minutes).Verify()
}

// Return HH:MM formatted time.
func (c Clock) String() string {
	h := c / 60
	m := c % 60

	return fmt.Sprintf("%02d:%02d", h, m)
}
