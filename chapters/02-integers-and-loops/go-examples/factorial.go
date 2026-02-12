package main

import "fmt"

func Factorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    result := Factorial(5)
    fmt.Println(result) // Вывод: 120
}