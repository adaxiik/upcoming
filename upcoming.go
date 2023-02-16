package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

type UpcomingMap map[string]time.Time

func upcomingAdd(filepath string, nameToAdd string, timeToAdd time.Time) error {
	upcoming, err := loadUpcoming(filepath)
	if err != nil {
		return err
	}

	if _, ok := upcoming[nameToAdd]; ok {
		return fmt.Errorf("name already exists")
	}

	upcoming[nameToAdd] = timeToAdd

	upcomingAutoRemove(upcoming)
	return saveUpcoming(upcoming, filepath)
}

func upcomingRemove(filepath string, nameToRemove string) error {
	upcoming, err := loadUpcoming(filepath)
	if err != nil {
		return err
	}

	if _, ok := upcoming[nameToRemove]; !ok {
		return fmt.Errorf("name does not exist")
	}

	delete(upcoming, nameToRemove)

	upcomingAutoRemove(upcoming)
	return saveUpcoming(upcoming, filepath)
}

func upcomingPrint(filepath string) error {

	upcoming, err := loadUpcoming(filepath)
	if err != nil {
		return err
	}

	if upcomingAutoRemove(upcoming) {
		err = saveUpcoming(upcoming, filepath)
		if err != nil {
			return err
		}
	}

	// sort by time
	keys := make([]string, 0, len(upcoming))
	for k := range upcoming {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return upcoming[keys[i]].Before(upcoming[keys[j]])
	})

	for _, key := range keys {
		timeleft := timeLeft(upcoming[key])
		fmt.Printf("%s: %s (%s)\n", key, timeLeftString(timeleft), timeleft)
	}
	return nil
}

// returns true if the upcoming map was changed
func upcomingAutoRemove(upcoming UpcomingMap) bool {
	changed := false
	for key, value := range upcoming {
		if timeLeft(value) <= 0 {
			delete(upcoming, key)
			changed = true
		}
	}
	return changed
}

func main() {

	mode := flag.String("mode", "print", "add, remove, print")
	filepath := flag.String("filepath", "", "filepath (always required)")
	timeToAdd := flag.String("time", "", "time to add (required in add mode)")
	name := flag.String("name", "", "name to add (required in add and remove mode)")

	flag.Parse()
	if !isFlagPassed("filepath") {
		fmt.Fprintln(os.Stderr, "filepath is required")
		os.Exit(1)
	}

	var err error
	switch *mode {
	case "add":
		if !isFlagPassed("time") || !isFlagPassed("name") {
			fmt.Fprintln(os.Stderr, "time and name are required")
			os.Exit(1)
		}

		upcomingTime, notOk := parseUpcomingTime(*timeToAdd)
		if notOk != nil {
			fmt.Fprintln(os.Stderr, notOk)
			os.Exit(1)
		}

		err = upcomingAdd(*filepath, *name, upcomingTime)
	case "remove":
		err = upcomingRemove(*filepath, *name)
	case "print":
		err = upcomingPrint(*filepath)
	default:
		fmt.Fprintln(os.Stderr, "invalid mode, using print as default")
		upcomingPrint(*filepath)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
