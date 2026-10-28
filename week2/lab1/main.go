
package main
import "fmt"
func main() {
    var a, b int

    fmt.Print("Enter the first integer (a): ")
    fmt.Scan(&a)
    fmt.Print("Enter the second integer (b): ")
    fmt.Scan(&b)


    // 1. Assignment operator (=)
    result := a
    result = b
  

    // 2. Addition assignment (+=)
    result = a
    result += b
    fmt.Println(a, "+=",b,":", result)

    // 3. Subtraction assignment (-=)
    result = a
    result -= b
    fmt.Println(a, "-=",b,":", result)

    // 4. Multiplication assignment (*=)
    result = a
    result *= b
    fmt.Println(a, "*=",b,":", result)

    // 5. Division assignment (/=)
    if b != 0 {
        result = a
        result /= b
        fmt.Println(a, "/=",b,":", result)

        // 6. Modulus assignment (%=)
        result = a
        result %= b
        fmt.Println(a, "+=",b,":"    , result)
    } else {
        fmt.Println("a /= b     : undefined (division by zero)")
        fmt.Println("a %= b     : undefined (modulo by zero)")
    }
}
