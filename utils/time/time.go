package time

import "time"

// GetDate return a date. eg 2006-01-02
func GetDate() string {

	return time.Now().Format("2006-01-02")
}

// GetTime return a time. eg 15:04:05.000000
func GetTime() string {
	return time.Now().Format("15:04:05.000000")
}

// GetDateTime return a datetime. eg 2006-01-02 15:04:05.000000
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}
