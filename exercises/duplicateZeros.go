package main

import "fmt"

func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            // Shift the elements
            for j := len(arr) -1 ; j > i; j-- {
                arr[j] = arr[j-1]
            }
            i++
        }
    }
    fmt.Println(arr)
}

func main() {
    duplicateZeros([]int{1,0,2,3,0,4,5,6})
}
