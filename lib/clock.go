package lib

import (
	"time"
)

// "Clock is an interface that has a method called Now that returns a time.Time".
// It exists to allow to override mocked versions on time for testing purposes
// For more info see https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-during-test
// @property Now - Returns the current time.
type Clock interface {
	Now() time.Time
}

// RealClock is a type that implements the Clock interface by calling the time.Now() function.
type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

func NewClock() *RealClock {
	return &RealClock{}
}
