package stealth

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/go-rod/rod"
)

func ScrollHumanLike(page *rod.Page, times int) {
    fmt.Println("Simulating human-like scrolling")

    if times < 1 {
        times = 1
    }

    for i := 0; i < times; i++ {
        distance := rand.Intn(400) + 200

        page.MustEval(`(d) => {
            window.scrollBy(0, d)
        }`, distance)

        pause := rand.Intn(800) + 300
        time.Sleep(time.Duration(pause) * time.Millisecond)
    }
}
