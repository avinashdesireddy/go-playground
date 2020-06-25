package main

import "fmt"

func sortedSquares(A []int) []int {
    // Square the numbers and store result back in A
    for k, v := range A {
        A[k] = v * v
    }

    for i := 1; i < len(A); i++ {
        for j := i; j > 0; j-- {
            if A[j] < A[j-1] {
                A[j-1], A[j] = A[j], A[j-1]
            }
        }
    }
    // Sort array
    return A
}

func main() {
    input := []int{-4,-1,0,3,10}
    fmt.Println(sortedSquares(input))
}
