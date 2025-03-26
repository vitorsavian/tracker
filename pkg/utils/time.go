package utils

import "time"

func UTCTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000000Z")
}
