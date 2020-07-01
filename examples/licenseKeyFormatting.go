package main

import "fmt"
import "strings"

func Reverse(s string) (result string) {
  for _, v := range s {
    result = string(v) + result
  }
  return
}
func main() {
  //S := "5F3-2e-9-Zw"
  S := "2-5G-3-j"
  K := 2
  var strBuffer strings.Builder
  var kCount int

  sFormatted := strings.ReplaceAll(S, "-", "")

  for i := len(sFormatted)-1; i >= 0; i-- {
    ch := rune(sFormatted[i])
    if kCount < K {
      strBuffer.WriteRune(ch)
      kCount += 1
    } else {
      kCount = 1
      strBuffer.WriteRune('-')
      strBuffer.WriteRune(ch)
    }
  }

  fmt.Println(strings.ToUpper(Reverse(strBuffer.String())))
}
