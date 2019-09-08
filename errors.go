package indent

import (
	"errors"
)

var (
	errInternalError  = errors.New("indent: Internal Error")
	errNilDestination = errors.New("indent: Nil Destination")
	errNilReceiver    = errors.New("indent: Nil Receiver")
	errNilSource      = errors.New("indent: Nil Source")
	errNilWriter      = errors.New("indent: Nil Writer")
	errUTF8RuneError  = errors.New("indent: UTF-8 Rune Error")
)
