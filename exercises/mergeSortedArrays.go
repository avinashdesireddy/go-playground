package main

import (
  "fmt"
)

func main() {
  arr1 := []int{3,6,7,9}
  arr2 := []int{1,2,4,6}

  var resultArr = []int{}

  i := 0
  j := 0

  for ; i < len(arr1) && j < len(arr2); {
    if arr1[i] <= arr2[j] {
      resultArr = append(resultArr, arr1[i])
      i += 1
    } else if arr1[i] > arr2[j] {
      resultArr = append(resultArr, arr2[j])
      j += 1
    }
  }
  if i != len(arr1) {
    resultArr = append(resultArr, arr1[i:]...)
  } else if j != len(arr2) {
    resultArr = append(resultArr, arr2[j:]...)
  }

  fmt.Println(resultArr)
}
