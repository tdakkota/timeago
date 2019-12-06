package timeago

import (
	"fmt"
	"math"
	"time"
)

type Template = string

type Locale interface {
	Get(DateAgoValues) Template
	GetLast(DateAgoValues) Template
}

var DefaultLocale = EnLocale{}

func stringForDuration(locale Locale, duration time.Duration) string {
	if duration.Hours() < 24 {
		if duration.Hours() >= 1 {
			return localizedStringForLocale(locale, HoursAgo, int(round(duration.Hours())))
		} else if duration.Minutes() >= 1 {
			return localizedStringForLocale(locale, MinutesAgo, int(round(duration.Minutes())))
		} else {
			return localizedStringForLocale(locale, SecondsAgo, int(round(duration.Seconds())))
		}
	} else {
		if duration.Hours() >= 8760 {
			years := duration.Hours() / 8760
			return localizedStringForLocale(locale, YearsAgo, int(years))
		} else if duration.Hours() >= 730 {
			months := duration.Hours() / 730
			return localizedStringForLocale(locale, MonthsAgo, int(months))
		} else if duration.Hours() >= 168 {
			weeks := duration.Hours() / 168
			return localizedStringForLocale(locale, WeeksAgo, int(weeks))
		} else {
			days := duration.Hours() / 24
			return localizedStringForLocale(locale, DaysAgo, int(days))
		}
	}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func localizedStringForLocale(locale Locale, valueType DateAgoValues, value int) string {
	if value >= 2 {
		return fmt.Sprintf(locale.Get(valueType), value)
	} else {
		return locale.GetLast(valueType)
	}
}
