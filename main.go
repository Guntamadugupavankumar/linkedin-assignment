package main

import (
    "fmt"
    "time"

    "linkedin-fresh/internal/auth"
    "linkedin-fresh/internal/search"
    "linkedin-fresh/internal/connect"
    "linkedin-fresh/internal/message"
    "linkedin-fresh/internal/config"
    "linkedin-fresh/internal/browser"
    "linkedin-fresh/internal/stealth"
)

func main() {
    fmt.Println("App started")

    cfg := config.LoadConfig()
    limiter := stealth.NewRateLimiter(cfg.DailyLimit)
    session := stealth.LoadSession()

    br, err := browser.Start()
    if err != nil {
        fmt.Println("Browser failed to start")
        return
    }
    defer br.Close()

    page, _ := br.OpenTestPage()
    stealth.ApplyFingerprint(page)
    stealth.RandomDelay(800, 2000)
    stealth.MoveMouseHumanLike(
        stealth.Point{X: 100, Y: 100},
        stealth.Point{X: 600, Y: 400},
        40,
    )

    stealth.RandomDelay(500, 1200)
    stealth.TypeHumanLike("Searching profiles")

    stealth.RandomDelay(500, 1200)
    stealth.ScrollHumanLike(page, 3)

    detection := stealth.CheckDetectionSignals(page)
    if detection.BotDetected {
        fmt.Println("Detection risk:", detection.Reason)
        fmt.Println("Stopping automation to stay safe")
        return
    }


    stealth.RandomDelay(500, 1500)
    if limiter.Allow("login") {
        auth.Login()
    }

    stealth.RandomDelay(800, 2000)
    if limiter.Allow("search") {
        search.SearchProfiles()
    }

    stealth.RandomDelay(800, 2000)
    if limiter.Allow("connect") {
        connect.SendConnections()
    }

    stealth.RandomDelay(800, 2000)
    if limiter.Allow("message") {
        message.SendMessages()
    }

    session.LastRun = time.Now()
    session.ActionsRun = limiter.Count
    stealth.SaveSession(session)

    fmt.Println("Daily limit:", cfg.DailyLimit)
    fmt.Println("App finished")
}



