// Parallel sort lists on multiple threads using fair scheduling
// With default params, 4 jobs get scheduled correctly on 4 threads

package main

import "fmt"
import "flag"
import "time"
import "math/rand"
import "sort"

var prng *rand.Rand

type ThreadState int
const (
    Unknown ThreadState = 0
    Idle    ThreadState = 1
    Busy    ThreadState = 2
    Paused  ThreadState = 3
    Done    ThreadState = 4
)

type ThreadMessage int
const (
    Pause   ThreadMessage = 0 + iota
    Resume
    Cancel
    End
)

func sortnums(n int) {
    seq := make([]int, n)

    randmax := n * 100
    for i := 0; i < n; i++ {
        seq[i] = prng.Intn(randmax)
    }
    sort.Ints(seq)
}

func worker(state chan<- ThreadState, dispatcher <-chan func(), i int) {
    fmt.Println("Worker", i, "started")
    state <- Idle
    for {
        job, more := <-dispatcher
        if more {
            fmt.Println("Worker", i, "starting job")
            state <- Busy
            start := time.Now()
            job()
            end := time.Now()
            state <- Idle
            fmt.Println("Worker", i, "completed job in", end.Sub(start))
        } else {
            fmt.Println("Worker", i, "finished all jobs")
            state <- Done
            break
        }
    }
    close(state)
}

func schedule(n int, l int, t int) {

    // Init
    states := make([]chan ThreadState, t)
    dispatchers := make([]chan func(), t)
    for i := 0; i < t; i++ {
        states[i] = make(chan ThreadState, 15)
        dispatchers[i] = make(chan func(), 15)

        fmt.Println("Scheduling worker", i)
        go worker(states[i], dispatchers[i], i)
    }

    knownStates := make([]ThreadState, t)
    remaining := l
    for {
        for i := 0; i < t; i++ {
            select {
            case knownStates[i] = <-states[i]:
                if knownStates[i] == Idle && remaining > 0 {
                    fmt.Println("Thread", i, "available for scheduling dispatch")
                    dispatchers[i] <- func () {
                        sortnums(n)
                    }
                    remaining--
                    fmt.Println("Scheduled on thread", i)
                    fmt.Println("Remaining jobs:", remaining)
                }
            default:
                continue
            }
        }
        if remaining == 0 {
            fmt.Println("Cleaning up")
            for i := 0; i < t; i++ {
                close(dispatchers[i])
                for {
                    _, more := <-states[i]
                    if !more {
                        fmt.Println("Thread", i, "terminated")
                        break
                    }
                }
            }
            return
        }
    }
}

func main() {
    prng = rand.New(rand.NewSource(time.Now().UnixNano()))

    var n int
    flag.IntVar(&n, "n", 10, "No. of integers per list")
    var l int
    flag.IntVar(&l, "l", 4, "No. of lists to sort")
    var t int
    flag.IntVar(&t, "t", 4, "No. of parallel threads")
    flag.Parse()

    schedule(n, l, t)

    fmt.Println("Done")
}
