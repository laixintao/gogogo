package main

import "fmt"

func sum(x, y int) (result int){
    result = x + y
    a := result
    return a
}

func main() {
    fmt.Println(sum(100, 300))
}
