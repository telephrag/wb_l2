package develop

import (
	"errors"
	"time"

	"github.com/beevik/ntp"
)

var (
	ErrLeapNotInSync = errors.New("leap second is not in sync on time server")
)

// Returns precise time as close as possible
// to the moment of return from the function bellow.
func GetPreciseTime() (time.Time, error) {

	leap := time.Second

	// Check if addition of leap second is required in advance
	// to neglect influence of execution of code bellow (it takes time)
	// on actuallity of received time.
	if leapCheck, err := ntp.Query("0.beevik-ntp.pool.ntp.org"); err == nil {
		switch {
		case leapCheck.Leap == ntp.LeapNoWarning:
			leap *= 0
		case leapCheck.Leap == ntp.LeapDelSecond:
			leap *= -1
		case leapCheck.Leap == ntp.LeapNotInSync:
			return time.Time{}, ErrLeapNotInSync
		}
	} else {
		return time.Time{}, err
	}

	result, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	// Accounting for time packet from the server was in flight
	if err == nil {
		return result.Time.Add(leap).Add(result.RTT), err
	}
	return time.Time{}, err
}
