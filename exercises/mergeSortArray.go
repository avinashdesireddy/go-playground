package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

    i,j := m-1,n-1
    if m == 0 {
        for k := 0; k < n; k++ {
            nums1[k] = nums2[k]
        }
    }
    for k := m+n-1; i >= 0 && j >= 0 ; k-- {
        if nums1[i] >= nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
    fmt.Println(nums1)
}
func main() {
    nums1 := []int{0}
    nums2 := []int{1}
    merge(nums1, 0, nums2, 1)
}
