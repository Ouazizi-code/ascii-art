package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . <text>")
        return
    }

    text := os.Args[1]
    banner := createBanner(text)
    fmt.Print(banner)
}

func createBanner(text string) string {
    lines := []string{
        " _    _          _   _          ",
        "| |  | |        | | | |         ",
        "| |__| |   ___  | | | |   ___   ",
        "|  __  |  / _ \\ | | | |  / _ \\  ",
        "| |  | | |  __/ | | | | | (_) | ",
        "|_|  |_|  \\___| |_| |_|  \\___/  ",
        "                                ",
        "                                ",
    }

    return strings.Join(lines, "\n") + "\n"
}
