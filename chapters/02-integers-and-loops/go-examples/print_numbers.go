package main

import "fmt"

func PrintNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    PrintNumbers()
}