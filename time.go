package main

import "time"

func getTimeInISOString() string {
	//returns this format: 2022-01-19T05:20:34.022Z
	return time.Now().UTC().Format("2006-01-02T15:04:05.006Z")
}
