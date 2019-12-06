// Package time provides a set of functions to return how much
// time has been passed between two dates.
package timeago

import (
	"time"
)

var DefaultTimeAgo = TimeAgo{DefaultLocale}

type TimeAgo struct {
	locale Locale
}

func NewTimeAgo(locale Locale) TimeAgo {
	return TimeAgo{locale: locale}
}

// FromNowWithTime takes a specific end Time value
// and the current Time to return how much has been passed
// between them.
func (ta TimeAgo) FromNowWithTime(end time.Time) (string, error) {
	return ta.WithTime(time.Now(), end)
}

// TimeAgoFromNowWithTime takes a specific end Time value
// and the current Time to return how much has been passed
// between them.
func TimeAgoFromNowWithTime(end time.Time) (string, error) {
	return DefaultTimeAgo.FromNowWithTime(end)
}

// FromNowWithString takes a specific layout as time
// format to parse the time string on end paramter to return
// how much time has been passed between the current time and
// the string representation of the time provided by user.
func (ta TimeAgo) FromNowWithString(layout, end string) (string, error) {
	t, e := time.Parse(layout, end)
	if e == nil {
		return ta.WithTime(time.Now(), t)
	} else {
		return "", ErrInvalidFormat
	}
}

// TimeAgoFromNowWithString takes a specific layout as time
// format to parse the time string on end paramter to return
// how much time has been passed between the current time and
// the string representation of the time provided by user.
func TimeAgoFromNowWithString(layout, end string) (string, error) {
	return DefaultTimeAgo.FromNowWithString(layout, end)
}

// WithTime takes a specific start/end Time values
// and calculate how much time has been passed between them.
func (ta TimeAgo) WithTime(start, end time.Time) (string, error) {
	duration := start.Sub(end)
	return stringForDuration(ta.locale, duration), nil
}

// TimeAgoWithTime takes a specific start/end Time values
// and calculate how much time has been passed between them.
func TimeAgoWithTime(start, end time.Time) (string, error) {
	return DefaultTimeAgo.WithTime(start, end)
}

// WithString takes a specific layout as time
// format to parse the time string on start/end parameter to return
// how much time has been passed between them.
func (ta TimeAgo) WithString(layout, start, end string) (string, error) {
	timeStart, e := time.Parse(layout, start)
	if e != nil {
		return "", ErrInvalidStartTimeFormat
	}

	timeEnd, e := time.Parse(layout, end)
	if e != nil {
		return "", ErrInvalidEndTimeFormat
	}

	duration := timeStart.Sub(timeEnd)
	return stringForDuration(ta.locale, duration), nil
}

// TimeAgoWithString takes a specific layout as time
// format to parse the time string on start/end parameter to return
// how much time has been passed between them.
func TimeAgoWithString(layout, start, end string) (string, error) {
	return DefaultTimeAgo.WithString(layout, start, end)
}
