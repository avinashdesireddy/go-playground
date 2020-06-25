package main

import "fmt"

func removeElement(nums []int, val int) []int {

    i := 0
    j := len(nums) -1
    for i != j {
        if nums[j] == val && nums[i] == val {
            j--
        } else if nums[i] != val {
            i++
        } else if nums[i] == val && nums[j] != val {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        } else if nums[i] != val && nums[j] == val {
            i++
            j--
        } else {
            break
        }
    }
    fmt.Println(j)
    return nums[:j]
}

func main() {
    fmt.Println(removeElement([]int{1,3,2}, 3))
}
