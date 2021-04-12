package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("tock")
	}
}

func main() {
	fmt.Println("Tickers tut - Go Jea")

	go backgroundTask()

	fmt.Println("resto de applicaccion continua")

	select {}

}
