package stealth

import (
    "fmt"
    "math/rand"
    "time"
)

func TypeHumanLike(text string) {
    fmt.Println("Simulating human-like typing")

    for _, ch := range text {
        fmt.Printf("%c", ch)

        delay := rand.Intn(120) + 40
        time.Sleep(time.Duration(delay) * time.Millisecond)
    }

    fmt.Println()
}
