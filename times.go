package golodash

import (
	"strconv"
	"strings"
	"time"
)

// FormatDuration formats a time range into a specified expression, such as MM:DD HH:mm:ss.
func FormatDuration(d time.Duration, formatStr string) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	milliseconds := int(d.Milliseconds()) % 1000

	replacements := map[string]string{
		"YYYY": "0000",
		"MM":   "00",
		"DD":   "00",
		"HH":   PadLeft(strconv.Itoa(hours), 2),
		"mm":   PadLeft(strconv.Itoa(minutes), 2),
		"ss":   PadLeft(strconv.Itoa(seconds), 2),
		"SS":   PadLeft(strconv.Itoa(milliseconds), 3),
	}

	for key, value := range replacements {
		formatStr = strings.ReplaceAll(formatStr, key, value)
	}
	return formatStr
}
