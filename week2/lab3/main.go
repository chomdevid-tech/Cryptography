
package main

import "fmt"

func main() {
    var a, b int

    // Get two numbers from the user
    fmt.Print("Enter first number: ")
    fmt.Scan(&a)

    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    // BITWISE OPERATORS
    fmt.Println("\n--- Bitwise Operators ---")

    fmt.Println("XOR (^):", a^b)  // Different bits give 1
    fmt.Println("NOT a (^a):", ^a) // Flip all bits of a
    fmt.Println("NOT b (^b):", ^b) // Flip all bits of b
    fmt.Println("OR (|):", a|b)   // Either bit is 1
    fmt.Println("AND (&):", a&b)  // Both bits must be 1

    // Shift bits left and right
    fmt.Println("a shifted left:", a<<1)
    fmt.Println("a shifted right:", a>>1)

    // ASSIGNMENT OPERATORS
    fmt.Println("\n--- Assignment Operators ---")

    result := a

    result += b
    fmt.Println("Addition (+=):", result)

    result = a
    result -= b
    fmt.Println("Subtraction (-=):", result)

    result = a
    result *= b
    fmt.Println("Multiplication (*=):", result)

    if b != 0 {
        result = a
        result /= b
        fmt.Println("Division (/=):", result)

        result = a
        result %= b
        fmt.Println("Remainder (%=):", result)
    } else {
        fmt.Println("Cannot divide by zero!")
    }
}
