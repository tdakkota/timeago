package timeago

type EnLocale struct{}

func (e EnLocale) Get(valueType DateAgoValues) Template {
	switch valueType {
	case YearsAgo:
		return "%d years ago"
	case MonthsAgo:
		return "%d months ago"
	case WeeksAgo:
		return "%d weeks ago"
	case DaysAgo:
		return "%d days ago"
	case HoursAgo:
		return "%d hours ago"
	case MinutesAgo:
		return "%d minutes ago"
	case SecondsAgo:
		return "%d seconds ago"
	}
	return ""
}

func (e EnLocale) GetLast(valueType DateAgoValues) Template {
	switch valueType {
	case YearsAgo:
		return "Last year"
	case MonthsAgo:
		return "Last month"
	case WeeksAgo:
		return "Last week"
	case DaysAgo:
		return "Yesterday"
	case HoursAgo:
		return "An hour ago"
	case MinutesAgo:
		return "A minute ago"
	case SecondsAgo:
		return "Just now"
	}
	return ""
}
