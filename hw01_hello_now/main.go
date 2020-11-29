package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
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
	fmt.Printf("exact time: %s\n", formatTime(exactTime))
}
