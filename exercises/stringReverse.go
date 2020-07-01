package main

import (
  "fmt"
)

func main() {
  str := "Hi, this is Avinash"
  charArray := []rune(str)
  i,j := 0, len(charArray) -1

  for ; i < j; {
    charArray[i], charArray[j] = charArray[j], charArray[i]
    i += 1
    j -= 1
  }
  fmt.Println(string(charArray))
}
