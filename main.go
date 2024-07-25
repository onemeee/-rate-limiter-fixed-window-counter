package main

import (
	"fmt"
	"time"
)

func main() {
	rateLimiter := NewRateLimiter(5, 10*time.Second)
	ip := "192.168.1.1"

	for i := 0; i < 7; i++ {
		if rateLimiter.Allow(ip) {
			fmt.Printf("Request %d from %s allowed.\n", i+1, ip)
		} else {
			fmt.Printf("Request %d from %s denied.\n", i+1, ip)
		}
		time.Sleep(2 * time.Second)
	}
}
