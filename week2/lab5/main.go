package main

import (
    "encoding/base64"
    "encoding/hex"
    "fmt"
)

func main() {
    var text string

    fmt.Print("Enter text: ")
    fmt.Scan(&text)

    // Binary
    fmt.Print("Binary: ")
    for i := 0; i < len(text); i++ {
        fmt.Printf("%08b ", text[i])
    }

    // Hexadecimal
    fmt.Println("\nHexadecimal:", hex.EncodeToString([]byte(text)))

    // Base64
    fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}

