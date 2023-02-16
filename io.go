package main

import (
	"encoding/json"
	"flag"
	"os"
)

func loadUpcoming(filepath string) (UpcomingMap, error) {

	file, err := openOrCreate(filepath, false)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	length, _ := file.Seek(0, 2)

	if length == 0 {
		return make(UpcomingMap), nil
	}

	file.Seek(0, 0)
	bytes := make([]byte, length)

	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}

	var upcoming UpcomingMap
	err = json.Unmarshal(bytes, &upcoming)

	return upcoming, err
}

func saveUpcoming(upcoming UpcomingMap, filepath string) error {
	bytes, err := json.Marshal(upcoming)

	if err != nil {
		return err
	}

	file, err := openOrCreate(filepath, true)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	return err
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func openOrCreate(filepath string, write bool) (*os.File, error) {
	if write {
		return os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	}
	return os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
}
