package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func formatTime(t time.Time) time.Time {
	return t.Round(0)
}

func main() {
	localTime := time.Now()
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("current time: %s\n", formatTime(localTime))
	fmt.Printf("exact time: %s", formatTime(exactTime))
}
