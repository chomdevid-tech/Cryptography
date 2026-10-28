
package main

import (
    "encoding/hex"
    "fmt"
)

func xorEncrypt(text string, key byte) string {
    result := ""

    for i := 0; i < len(text); i++ {
        result += string(text[i] ^ key)
    }

    return result
}

func main() {
    var text, keyInput string

    fmt.Print("Enter plaintext (one word): ")
    fmt.Scan(&text)

    fmt.Print("Enter one-character key: ")
    fmt.Scan(&keyInput)

    key := keyInput[0]

    // Encrypt
    encrypted := xorEncrypt(text, key)

    // Decrypt using the same function and key
    decrypted := xorEncrypt(encrypted, key)

    fmt.Println("Plaintext:", text)
    fmt.Println("Encrypted (hex):", hex.EncodeToString([]byte(encrypted)))
    fmt.Println("Decrypted:", decrypted)
}
