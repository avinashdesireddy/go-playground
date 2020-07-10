package main

import "fmt"

func main() {
  // Bubble sort concept is to compare 2 elements
  // and move the larges to the end of the array

  arr := [10]int{3,8,6,3,2,5,6,3,6,9}
  for i := 0; i < len(arr); i++ {
    for j := i; j < len(arr); j++ {
      if arr[i] > arr[j] {
        // Swap numbers
        arr[i], arr[j] = arr[j], arr[i]
      }
    }
  }
  fmt.Println(arr)
}
