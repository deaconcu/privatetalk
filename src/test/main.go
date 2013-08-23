package test

import (
    "config"
    "fmt"
)

func main() {
    config.LoadConfig()
    fmt.Println("Hello, GO!")
}
