package main

import "fmt"
import "strings"

func main() {
  input := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com"}
  result := make(map[string]bool)

  for _,v := range input {
    emailParts := strings.Split(v, "@")
    // Valid Email -- emails[i] contains exactly one '@'
    if len(emailParts) == 2 {
      //fmt.Println("Vaid")
      // Ignore +
      localNameWithDot := strings.Split(emailParts[0], "+")
      local_name := strings.Replace(localNameWithDot[0], ".", "", -1)
      validEmail := local_name + "@" + emailParts[1]
      result[validEmail] = true
    }
  }
  fmt.Println(len(result))
}
