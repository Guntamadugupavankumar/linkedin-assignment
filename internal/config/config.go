package config

import "fmt"

type Config struct {
DailyLimit int
}

func LoadConfig() Config {
fmt.Println("Config loaded")
return Config{
DailyLimit: 5,
}
}