package main

import (
  "fmt"
)

func main() {
  arr := []int{2,4,5,6,3,4,2,3}
  // ans - 2
  hashMap := make(map[int]bool)
  for _, v := range arr {
    if !hashMap[v] {
      hashMap[v] = true
    } else {
      fmt.Println(v)
      break
    }
  }
}
