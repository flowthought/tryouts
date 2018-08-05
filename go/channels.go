// Parallel sort lists on multiple threads using fair scheduling

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
    Done    ThreadState = 3
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

func worker(state chan<- ThreadState, dispatcher <-chan func(), message <-chan ThreadMessage) {
    state <- Idle
    for {
        select {
        case job := <-dispatcher:
            state <- Busy
            job()
            state <- Idle
        case msg := <-message:
            switch msg {
            case Pause:
                // Implementation
            case Resume:
                // Implementation
            case Cancel:
                // Implementation
            case End:
                state <- Done
                close(state)
                return
            }
        default:
        }
    }
}

func schedule(n int, l int, t int) {

    // Init
    states := make([]chan ThreadState, t)
    dispatchers := make([]chan func(), t)
    messages := make([]chan ThreadMessage, t)
    for i := 0; i < t; i++ {
        states[i] = make(chan ThreadState, 5)
        dispatchers[i] = make(chan func(), 5)
        messages[i] = make(chan ThreadMessage, 5)

        go worker(states[i], dispatchers[i], messages[i])
    }

    knownStates := make([]ThreadState, t)
    remaining := l
    for {
        for i := 0; i < t; i++ {
            select {
            case knownStates[i] = <-states[i]:
                if knownStates[i] == Idle && remaining > 0 {
                    dispatchers[i] <- func () {
                        sortnums(n)
                    }
                    remaining--
                }
            default:
                continue
            }
        }
        if remaining == 0 {
            for i := 0; i < t; i++ {
                messages[i] <- End
                close(messages[i])
                close(dispatchers[i])
            }
            break
        }
    }

    return
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
