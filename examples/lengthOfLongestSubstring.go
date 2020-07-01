package main

import (
  "fmt"
)

func lengthOfLongestSubstring(s string) int {
  // define hashset
  uniqSet := make(map[byte]bool)
  var max int

  a_pointer, b_pointer := 0, 0

  for ;b_pointer < len(s); {
    // Insert char to set of not exists and increment max
    if !uniqSet[s[b_pointer]] {
      uniqSet[s[b_pointer]] = true
      b_pointer += 1
      if len(uniqSet) > max {
        max = len(uniqSet)
      }
    } else {
      delete(uniqSet, s[a_pointer])
      a_pointer += 1
    }
  }
  fmt.Println(uniqSet)
  return max
}

func main() {
  fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
