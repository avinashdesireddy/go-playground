package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    max,cMax := 0,0

    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            cMax = 0
            continue
        }
        if nums[i] == 1 {
            cMax += 1
        }

        if max < cMax {
            max = cMax
        }
    }

    return max
}

func main() {
//    a := []int{1,1,0,1,1,1}
//    a := []int{1,1,0,0,0,0,0,1,1}
    a := []int{0}
    fmt.Println(maxConsecutive(a))
}
