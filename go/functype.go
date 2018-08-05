package main

import "fmt"

func foo() {
    fmt.Println("This is a function")
}

func main() {
    fmt.Printf("Function type is %T\n", foo)

    var f func()
    f = foo
    f()
}
