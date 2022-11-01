package develop

import "fmt"

type UnpackError struct {
	Err      error
	Location int
}

func (e UnpackError) Error() string {
	return fmt.Sprintf("%s: %d", e.Err, e.Location)
}

func (e UnpackError) Unwrap() error {
	return e.Err
}
