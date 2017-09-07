package downlink

import "errors"

// downlink errors
var (
	ErrFPortMustNotBeZero     = errors.New("FPort must not be 0")
	ErrFPortMustBeZero        = errors.New("FPort must be 0")
	ErrInvalidAppFPort        = errors.New("FPort must be between 1 and 224")
	ErrNoLastRXInfoSet        = errors.New("no last RX-Info set available")
	ErrInvalidDataRate        = errors.New("invalid data-rate")
	ErrMaxPayloadSizeExceeded = errors.New("maximum payload size exceeded")
	ErrAbort                  = errors.New("nothing to do")
)
