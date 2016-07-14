package gpio

import "fmt"

type UnsupportedModeError struct {
	mode Mode
}

func (e *UnsupportedModeError) Error() string {
	return fmt.Sprintf("gpio: unsupported GPIO pin mode %s", e.mode)
}
