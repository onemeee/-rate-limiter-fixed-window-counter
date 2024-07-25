package main

import (
    "time"
)

type Request struct {
    IP       string
    Count    int
    LastSeen time.Time
}

func (r *Request) Cleanup(now time.Time, window time.Duration) {
    if now.Sub(r.LastSeen) >= window {
        r.Count = 0
    }
    r.LastSeen = now
}