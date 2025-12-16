package browser

import (
    "fmt"
    "time"

    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
    Rod *rod.Browser
}

func Start() (*Browser, error) {
    fmt.Println("Starting browser")

    path := launcher.New().
        Leakless(false). 
        Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe").
        Headless(true).
        MustLaunch()

    b := rod.New().ControlURL(path).MustConnect()

    return &Browser{
        Rod: b,
    }, nil
}

func (b *Browser) OpenTestPage() (*rod.Page, error) {
    fmt.Println("Opening test page")

    page := b.Rod.MustPage("https://example.com")
    page.MustWaitLoad()

    time.Sleep(2 * time.Second)

    return page, nil
}


func (b *Browser) Close() {
    fmt.Println("Closing browser")
    b.Rod.MustClose()
}
