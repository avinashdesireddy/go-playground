package main

import "fmt"

func mergeSort(arr []int) []int {
  if len(arr) <= 1 {
    return arr
  }

  AB := merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))

  return AB
}

func merge(A []int, B []int) []int {
  AB := []int{}

  i := 0
  j := 0
  for ;i < len(A) && j < len(B); {
    if A[i] < B[j] {
      AB = append(AB, A[i])
      i++
    } else {
      AB = append(AB, B[j])
      j++
    }
  }

  if i != len(A) {
    AB = append(AB, A[i:]...)
  }
  if j != len(B) {
    AB = append(AB, B[j:]...)
  }

  return AB
}

func main() {
  // Merge sort using recurssive method
  arr := []int{3,8,6,3,2,5,6,3,6,9}
  fmt.Println("Array: ", arr)
  fmt.Println("Sorted: ", mergeSort(arr))
}
