package main

import (
	"fmt"
	"time"
)

func getTimeInISOString() string {
	//returns this format: 2022-01-19T05:20:34.022Z
	return time.Now().UTC().Format("2006-01-02T15:04:05.006Z")
}

func get30SecondsFromTime(t string) string {
	layout := "2006-01-02T15:04:05.000Z"
	tdd, err := time.Parse(layout, t)

	if err != nil {
		fmt.Println(err)
	}
	return string(tdd.Add(time.Second * 30).Format(layout))
}
