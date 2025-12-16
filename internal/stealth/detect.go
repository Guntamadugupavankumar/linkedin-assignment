package stealth

import (
    "fmt"

    "github.com/go-rod/rod"
)

type DetectionResult struct {
    BotDetected bool
    Reason      string
}

func CheckDetectionSignals(page *rod.Page) DetectionResult {
    fmt.Println("Running detection awareness checks")

    detected := page.MustEval(`() => {
        if (navigator.webdriver === true) return true
        if (!navigator.plugins || navigator.plugins.length === 0) return true
        return false
    }`).Bool()

    if detected {
        return DetectionResult{
            BotDetected: true,
            Reason:      "Suspicious browser signals detected",
        }
    }

    return DetectionResult{
        BotDetected: false,
        Reason:      "No obvious automation signals",
    }
}
