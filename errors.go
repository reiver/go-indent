package indent

import (
	"errors"
)

var (
	errNilReceiver = errors.New("indent: Nil Receiver")
	errNilWriter   = errors.New("indent: Nil Writer")
)
