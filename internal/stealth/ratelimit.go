package stealth

import "fmt"

type RateLimiter struct {
    Max   int
    Count int
}

func NewRateLimiter(max int) *RateLimiter {
    return &RateLimiter{
        Max:   max,
        Count: 0,
    }
}

func (r *RateLimiter) Allow(action string) bool {
    if r.Count >= r.Max {
        fmt.Println("Rate limit reached. Stopping actions.")
        return false
    }

    r.Count++
    fmt.Println("Allowed action:", action, "| Used:", r.Count, "/", r.Max)
    return true
}
