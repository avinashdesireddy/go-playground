package main

import "fmt"

type Node struct {
  data int
  next *Node
}

// First node -- Called "head"
var head *Node = nil

func main() {
  myList := new(Node)

  for head != nil {
    fmt.Printf("%d ", head.data)
    head = head.next
  }
}
