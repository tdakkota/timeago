package timeago

import "errors"

var (
	ErrInvalidFormat          = errors.New("invalid format")
	ErrInvalidStartTimeFormat = errors.New("invalid start time format")
	ErrInvalidEndTimeFormat   = errors.New("invalid end time format")
)
