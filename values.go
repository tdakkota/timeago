package timeago

type DateAgoValues int

const (
	SecondsAgo DateAgoValues = iota
	MinutesAgo
	HoursAgo
	DaysAgo
	WeeksAgo
	MonthsAgo
	YearsAgo
)
