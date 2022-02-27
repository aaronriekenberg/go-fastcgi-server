package utils

import "time"

const (
	CacheControlHeaderKey      = "cache-control"
	MaxAgeZero                 = "max-age=0"
	ContentTypeHeaderKey       = "content-type"
	ContentTypeApplicationJSON = "application/json"
)

func FormatTime(t time.Time) string {
	const timeFormat = "Mon Jan 2 15:04:05.000000000 -0700 MST 2006"

	return t.Format(timeFormat)
}