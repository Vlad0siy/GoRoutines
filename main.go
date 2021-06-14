package main

import (
  "fmt"
  "time"
  "math/rand"
)

type TDone struct {
  state bool
  routine int
}

func Routine(n int, ADone chan TDone) {
  for i := 0; i < 10; i++ {
    fmt.Println("Routine: ", n, ":", i)
    rnd := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * rnd)
  }
  ADone <- TDone{true, n}
}

func CheckDone(ADone chan TDone) {
  fmt.Println("checking...")
  for {
    typeDone := <-ADone
    fmt.Println("Routine: ", typeDone.routine, "done: ", typeDone.state)
    if typeDone.state {
      break
    }
  }
}

func main() {
  channel := make(chan TDone)
  for i := 0; i < 5; i++ {
      go Routine(i, channel)
      go CheckDone(channel)
  }
  var input string
  fmt.Scanln(&input)
}
