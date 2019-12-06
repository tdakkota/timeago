package timeago

type DateAgoValues int

const (
	JustNow DateAgoValues = iota
	SecondsAgo
	LastMinute
	MinutesAgo
	LastHour
	HoursAgo
	LastDay
	DaysAgo
	LastWeek
	WeeksAgo
	LastMonth
	MonthsAgo
	LastYear
	YearsAgo
)
