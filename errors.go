package lesphina

import (
	"errors"
)

var (
	ErrNoSuchFile      = errors.New("No Such File")
	ErrDirNotSupported = errors.New("Dir Not Supported")
)
