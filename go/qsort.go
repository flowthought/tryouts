// Partition Variant: https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme

package main

import "fmt"
import "time"
import "math/rand"
import "flag"

func main() {
    var size int
    flag.IntVar(&size, "n", 10, "No. of integers")
    flag.Parse()

    seq := make([]int, size)
    prng := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < len(seq); i++ {
        seq[i] = prng.Intn(100)
    }
    fmt.Println(seq)

    quicksort(seq, 0, len(seq) - 1)
    fmt.Println(seq)
}

func quicksort(seq []int, lo int, hi int) {
    if (hi > lo) {
        boundary := partition(seq, lo, hi)
        quicksort(seq, lo, boundary - 1)
        quicksort(seq, boundary + 1, hi)
    }
}

func partition(seq []int, lo int, hi int) int {
    pivot := seq[hi]
    boundary := lo - 1
    for i := lo; i <= hi; i++ {
        if seq[i] < pivot {
            boundary = boundary + 1
            seq[boundary], seq[i] = seq[i], seq[boundary]
        }
    }
    boundary = boundary + 1
    seq[boundary], seq[hi] = pivot, seq[boundary]
    return boundary
}
