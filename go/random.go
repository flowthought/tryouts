package main

import "fmt"
import "time"
import "math/rand"
import "flag"

func main() {
    var size int
    flag.IntVar(&size, "n", 5, "No. of integers")
    flag.Parse()

    seq := make([]int, size)
    prng := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < len(seq); i++ {
        seq[i] = prng.Intn(100)
    }
    fmt.Println(seq)
}
