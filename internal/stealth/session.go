package stealth

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type SessionState struct {
    LastRun    time.Time `json:"last_run"`
    ActionsRun int       `json:"actions_run"`
}

const sessionFile = "session.json"

func LoadSession() *SessionState {
    file, err := os.Open(sessionFile)
    if err != nil {
        fmt.Println("No previous session found")
        return &SessionState{}
    }
    defer file.Close()

    var state SessionState
    json.NewDecoder(file).Decode(&state)

    fmt.Println("Session loaded")
    return &state
}

func SaveSession(state *SessionState) {
    file, err := os.Create(sessionFile)
    if err != nil {
        fmt.Println("Failed to save session")
        return
    }
    defer file.Close()

    json.NewEncoder(file).Encode(state)
    fmt.Println("Session saved")
}
