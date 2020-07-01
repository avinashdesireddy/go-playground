package main

import (
  "fmt"
)

func main() {
    m := make(map[string]int)
    m["k1"] = 7

    v, ok := m["k1"]
    fmt.Println("ok?: ",ok, v)
}
