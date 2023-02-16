package main

import (
	"fmt"
	"strings"
	"time"
)

func parseUpcomingTime(input string) (time.Time, error) {

	// relative time
	if strings.HasPrefix(input, "+") {
		duration, err := time.ParseDuration(input[1:])
		if err != nil {
			return time.Time{}, err
		}
		return time.Now().Add(duration), nil
	}

	// smart formats
	smartTimeFormatsYear := []string{
		"02.01",
		"2.1"}

	for _, format := range smartTimeFormatsYear {
		t, err := time.Parse(format, input)
		if err == nil {
			return time.Date(time.Now().Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local), nil
		}
	}

	smartTimeFormatsDay := []string{
		"15:04",
		"15:04:05"}

	for _, format := range smartTimeFormatsDay {
		t, err := time.Parse(format, input)
		if err == nil {
			return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
		}
	}

	// full formats
	timeFormats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"02.01.2006 15:04:05",
		"02.01.2006 15:04",
		"02.01.2006",
		"02.01.06 15:04:05",
		"02.01.06 15:04",
		"02.01.06",
		"2.1.2006 15:04:05",
		"2.1.2006 15:04",
		"2.1.2006",
		"2.1.6 15:04:05",
		"2.1.6 15:04",
		"2.1.6"}

	for _, format := range timeFormats {
		t, err := time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format")

}
