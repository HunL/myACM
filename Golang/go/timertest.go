package main

import (
    "fmt"
    "time"
)

func testTimer1() {
    go func() {
        fmt.Println("test timer1")
    }()
}

func testTimer2() {
    go func() {
        fmt.Println("test timer2")
    }()
}

func timer1() {
    // timer1 := time.NewTicker(4 * time.Second)
    timer1 := time.NewTimer(4 * time.Second)
    select {
    case <-timer1.C:
        testTimer1()
    }
}

func timer2() {
    // timer2 := time.NewTicker(2 * time.Second)
    timer2 := time.NewTimer(2 * time.Second)
    select {
    case <-timer2.C:
        testTimer2()
    }
}

func main() {
    go timer1()
    go timer2()
    fmt.Println("will end")
    time.Sleep(5 * time.Second)
}
