package stealth

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandomDelay(minMs int, maxMs int) {
    if maxMs <= minMs {
        maxMs = minMs + 500
    }

    delay := rand.Intn(maxMs-minMs) + minMs
    time.Sleep(time.Duration(delay) * time.Millisecond)
}
