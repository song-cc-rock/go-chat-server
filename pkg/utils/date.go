package utils

import "time"

// GetNowTimeUnix get current time unix in milliseconds
func GetNowTimeUnix() int64 {
	return time.Now().UnixMilli()
}
