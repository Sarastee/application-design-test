package date

import "time"

// DaysBetween ...
func DaysBetween(from time.Time, to time.Time) []time.Time {
	fromDay := format(from)
	toDay := format(to)

	if fromDay.After(toDay) {
		return nil
	}

	days := []time.Time{}
	for d := fromDay; !d.After(toDay); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func format(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Date ...
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
