package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]*Request
	limit    int
	window   time.Duration
	mutex    sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]*Request),
		limit:    limit,
		window:   window,
		mutex:    sync.Mutex{},
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	req, exists := rl.requests[ip]
	if exists {
		req.Cleanup(now, rl.window)
		if req.Count < rl.limit {
			req.Count++
			req.LastSeen = now
			return true
		}
		return false
	}

	rl.requests[ip] = &Request{
		IP:       ip,
		Count:    1,
		LastSeen: now,
	}
	return true
}
