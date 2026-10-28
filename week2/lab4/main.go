
package main

import "fmt"

func add(a, b int) int {
    return a + b
}
func sub(a, b int) int {
    return a - b
}
func mul(a, b int) int {
    return a * b
}

func div(a, b int) int {
    return a / b
}

func mod(a, b int) int {
    return a % b
}

func main() {
    var a, b, choice int

    fmt.Println("===== Mini Calculator =====")
    fmt.Println("1. Add")
    fmt.Println("2. Subtract")
    fmt.Println("3. Multiply")
    fmt.Println("4. Divide")
    fmt.Println("5. Modulus")
    fmt.Print("Choose: ")
    fmt.Scan(&choice)

    fmt.Print("Enter a: ")
    fmt.Scan(&a)

    fmt.Print("Enter b: ")
    fmt.Scan(&b)

    switch choice {
    case 1:
        fmt.Println("Result:", add(a, b))
    case 2:
        fmt.Println("Result:", sub(a, b))
    case 3:
        fmt.Println("Result:", mul(a, b))
    case 4:
        if b == 0 {
            fmt.Println("Error: division by zero")
        } else {
            fmt.Println("Result:", div(a, b))
        }
    case 5:
        if b == 0 {
            fmt.Println("Error: modulo by zero")
        } else {
            fmt.Println("Result:", mod(a, b))
        }
    default:
        fmt.Println("Invalid choice")
    }
}


