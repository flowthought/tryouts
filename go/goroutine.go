// https://gobyexample.com/goroutines

package main

import "fmt"
import "flag"
import "time"
import "math/rand"

var prng *rand.Rand

func dowork(source string, count int) {
    for i := 1; i <= count; i++ {
        fmt.Println(source, ":", i)
        pause := prng.Intn(100)
        time.Sleep(time.Duration(pause) * time.Millisecond)
    }
}

func main() {
    var jobsize int
    flag.IntVar(&jobsize, "n", 5, "Size of the job")
    flag.Parse()

    prng = rand.New(rand.NewSource(time.Now().UnixNano()))

    dowork("main_thread", jobsize)

    go dowork("worker_thread", jobsize)

    go func(jobsize int) {
        dowork("anonymous_function", jobsize)
    }(jobsize)

    fmt.Scanln()
    fmt.Println("Done")
}
