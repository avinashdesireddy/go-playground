package main

import (
    "fmt"
    )

func linearSearch(dataList []int, key int) bool {
    for _, item := range dataList {
        if item == key {
            return true
        }
    }
    return false
}

func main() {
    items := []int{95, 78, 1,65, 7, 54}
    fmt.Println(linearSearch(items, 58))
}
