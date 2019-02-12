package main

import "fmt"
import "time"

func main() {
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }

    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("Weekend")
        default:
            fmt.Println("Weekday")
    }

    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("Before Noon")
        default:
            fmt.Println("After noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
            case bool:
                fmt.Println("I'm a Bool")
            case int:
                fmt.Println("I'm an Int")
            default:
                fmt.Printf("Don't know type %T\n", t)
            }
        }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
