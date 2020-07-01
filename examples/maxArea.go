package main

import "fmt"

func maxArea(height []int) int {
  var maxArea int
  var area int
  a_pointer, b_pointer := 0, len(height) -1

  for ;a_pointer < b_pointer; {
    if height[a_pointer] < height[b_pointer] {
      // a_pointer < b_pointer
      area = height[a_pointer] * (b_pointer - a_pointer)
      if area > maxArea {
        maxArea = area
      }
      a_pointer += 1
    } else {
      area = height[b_pointer] * (b_pointer - a_pointer)
      if area > maxArea {
        maxArea = area
      }
      b_pointer -= 1
    }
  }
  return maxArea
}

func main() {
  height := []int{1,8,6,2,5,4,8,3,7}
  fmt.Println(maxArea(height))
}
