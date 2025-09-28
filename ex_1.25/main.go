package main

import (
	"log"
	"time"
)

func sleep(d time.Duration) {
	select {
	case <-time.After(d):
	}
}

func CallSleep() {
	log.Println("Start sleep")
	sleep(5 * time.Second)
	log.Println("End sleep")
}

func main() {
	CallSleep()
}
