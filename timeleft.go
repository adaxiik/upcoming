package main

import (
	"fmt"
	"time"
)

func timeLeft(t time.Time) time.Duration {
	return time.Until(t).Round(time.Second)
}
func timeLeftString(t time.Duration) string {
	seconds := int(t.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours%24, minutes%60)
	}

	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes%60)
	}

	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds%60)
	}

	return fmt.Sprintf("%ds", seconds)
}
