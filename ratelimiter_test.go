package main

import (
    "testing"
    "time"
)

func TestRateLimiter(t *testing.T) {
    // 제한: 10초에 5개의 요청
    rateLimiter := NewRateLimiter(5, 10*time.Second)
    ip := "192.168.1.1"

    // 5개의 요청이 허용되어야 함
    for i := 0; i < 5; i++ {
        if !rateLimiter.Allow(ip) {
            t.Errorf("Request %d from %s should have been allowed", i+1, ip)
        }
    }

    // 6번째 요청은 거부되어야 함
    if rateLimiter.Allow(ip) {
        t.Errorf("Request 6 from %s should have been denied", ip)
    }

    // 10초 후에 요청이 다시 허용되어야 함
    time.Sleep(10 * time.Second)
    if !rateLimiter.Allow(ip) {
        t.Errorf("Request after reset from %s should have been allowed", ip)
    }
}

func TestRateLimiterWindowReset(t *testing.T) {
    // 제한: 5초에 2개의 요청
    rateLimiter := NewRateLimiter(2, 5*time.Second)
    ip := "192.168.1.2"

    // 2개의 요청이 허용되어야 함
    for i := 0; i < 2; i++ {
        if !rateLimiter.Allow(ip) {
            t.Errorf("Request %d from %s should have been allowed", i+1, ip)
        }
    }

    // 3번째 요청은 거부되어야 함
    if rateLimiter.Allow(ip) {
        t.Errorf("Request 3 from %s should have been denied", ip)
    }

    // 5초 후에 요청이 다시 허용되어야 함
    time.Sleep(5 * time.Second)
    if !rateLimiter.Allow(ip) {
        t.Errorf("Request after reset from %s should have been allowed", ip)
    }
}
