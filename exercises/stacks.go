package main

import "fmt"

// Stacks implementation using Arrays
var A []int

func push(a int) {
  A = append(A, a)
}

func pop() {
  fmt.Println("pop")
}

func printA() {
  for _, v := range A {
    fmt.Println(v)
  }
}

func main() {
  push(1)
  push(2)
  push(3)
  printA()
  pop()
  printA()
}
