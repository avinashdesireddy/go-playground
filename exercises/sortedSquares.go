package main

import "fmt"

func sortedSquares(A []int) []int {
    var output []int
    // Square the numbers and store result back in A
    for k, v := range A {
        A[k] = v * v
    }

    for i := 0; i < len(A); i++ {
        for j := i + 1; j < len(A) - 1; j++ {
            if A[i] > A[j] {
                
            }
        }
    }
    // Sort array
    return output
}

func main() {
    input := []int{-4,-1,0,3,10}
    fmt.Println(sortedSquares(input))
}
