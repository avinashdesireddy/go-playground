package main

import (
    "fmt"
    "os"
    "strconv"
    )

func main() {
    fmt.Println(len(os.Args))
    if len(os.Args) < 4 {
        fmt.Println("Required number of Args is 3, but provided ", len(os.Args))
        os.Exit(1)
    }
    var c byte = 'H'
    fmt.Println(c)
    mj := string(45)
    fmt.Println(mj)

    // Go represents symbols as numbers
    // and strings as arays of symbols
    // When Go sees the sting "Hello World" it visualizes it 
    // as an array, [72 101 108 108 111 ...]
    // the correspondence between symbols and numbers is called "ASCII"
    // String is array of bytes in Go

    // os.Args stores aln array of all command line parameters 
    // as an array of strings

    x, errx := strconv.Atoi(os.Args[1])
    if errx != nil { // There was a problem
        fmt.Println("Error converting first command line parameter!")
        os.Exit(1)
    }
    // if you make it here, all was ok with converting x
    y, erry := strconv.Atoi(os.Args[2])
    if erry != nil { // There was a problem
        fmt.Println("Error converting first command line parameter!")
        os.Exit(1)
    }

    fmt.Println(x + y)

    //strconv allows us to convert between types involving strings naturally

    mj = strconv.Itoa(45)
    fmt.Println(mj)

    z, errz := strconv.ParseFloat(os.Args[3], 64)
    if errz != nil {
        fmt.Println("Error converting third command line parameter")
        os.Exit(1)
    }
    fmt.Println("Third command line parameter is", z)
}
