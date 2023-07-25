# Session 3

## What is concurrency

concurrency is the ability for a program to be decomposed into parts that can run independently of each other.

This means that tasks can be executed out of order and the result would still be the same as if they are executed in order.

## Concurrency in go

Goroutine → A Goroutine is a function or method which executes independently and 
simultaneously in connection with any other Goroutines present in your 
program. Or in other words, every concurrently executing activity in Go 
language is known as a [Goroutines](https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/).

## goroutine vs thread

| Goroutines are managed by the go runtime. | Operating system threads are managed by kernel. |
| --- | --- |
| They are cooperatively scheduled. | They are preemptively scheduled. |
| Goroutines have easy communication medium known as channel. | Thread does not have easy communication medium. |

## Define goroutine in Go (`go` keyword)

```go
package main

func printer(str string) {
	for i := 0; i < 10; i++ {
		println(str)
	}
}
func main() {
	go printer("hello")
	printer("bye")

}
```

## time to code

write a program that takes a number from input and counts from 1 to infinitely and visa-versa simultaneously. finally print the max number and min number

### class work

implement the scenario discussed in class

## channels (buffered - unbuffered)

### unbuffered channels

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

<aside>
💡 By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

</aside>

### buffered channels

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

<aside>
💡 Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

</aside>

## range and close channel

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

## select

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received from channel 1", msg1)
        case msg2 := <-c2:
            fmt.Println("received from channel 2", msg2)
        }
    }
}
```