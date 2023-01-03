package common

import (
	"strings"
	"time"
)

// DateUS returns a date in format "Jan 2, 2006"
func DateUS(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format("Jan _2, 2006")
}

// DateUS returns a date in format "Jan 2, 2006 15:04:05"
func DateTimeUS(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format("Jan _2, 2006 15:04:05")
}

// DateISO returns a date in format "2006-01-02"
func DateISO(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format("2006-01-02")
}

// DateTimeISO returns a date in format "2006-01-02T15:04:05Z"
func DateTimeISO(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format(time.RFC3339)
}

// UnixToISO converts a Unix timestamp to ISO format.
func UnixToISO(ts int64) string {
	return time.Unix(ts, 0).Format(time.RFC3339)
}

// YesNo returns "Yes" if value is true, "No" if value is false.
func YesNo(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

// DefaultString returns value if it's non-empty.
// Otherwise, it returns _default.
func DefaultString(value, _default string) string {
	if len(strings.TrimSpace(value)) > 0 {
		return value
	}
	return _default
}
