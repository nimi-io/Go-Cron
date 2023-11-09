package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Africa/Lagos")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	cronJob := cron.NewWithLocation(loc)
	num := 0
	cronJob.AddFunc("*/2 * * * * *", func() {
		num++
		log.Println(num)

		log.Println("Cron job executed")
	})

	cronJob.Start()

	defer cronJob.Stop()

	for {
		select {
		case <-time.After(10 * time.Second):
			log.Println("Custom action after 10 seconds")

		case <-time.After(15 * time.Second):
			log.Println("Custom action after 15 seconds")
		}
	}

}
