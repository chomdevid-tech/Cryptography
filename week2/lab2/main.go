
package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Enter the first integer (a): ")
    fmt.Scan(&a)

    fmt.Print("Enter the second integer (b): ")
    fmt.Scan(&b)

    // 1. both positive
    bothPositive := a > 0 && b > 0
    fmt.Println("Both numbers are positive (&&):", bothPositive)

    // 2. one greater than other one 
    oneGreater := a > b || b > a
    fmt.Println("One number is greater than the other (||):", oneGreater)

    // 3. not equal 
    notEqual := !(a == b)
    fmt.Println("Numbers are not equal (!):", notEqual)
}
