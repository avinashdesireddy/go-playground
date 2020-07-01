## Go Playground
*********************

Go is an open source programming language designed for building simple, fast, and reliable software.

#### This repository contains code snippets as I learn Go
```Ref: https://gobyexample.com/```

Installing Go in OSX
```brew install golang```


### Data Structures in Golang

Basic data structures in Golang include
- Array
- Slice
- Map
- Struct

#### Array
Used to store fixed number of elements of same type. Once an array is defined, elements cannot be added or removed from the array.

Creating an array
```
var arrName [size]type
```

#### Slice
An Array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

##### Creating a  slice
```
var s []int --> Using var
s := []int{1,2,3,4,5} --> Shorthand notation
s := make([]int, n) --> n is the number of the elements
```

CURD operations on slice `s = []int{10,20,30,40}`

1. Add to a slice
   ```
   s = append(s, 50) // Adding one element
   s = append(s, 60, 70) // Adding more than one element
   ```
2. Delete from a slice
   ```
   // Inserting element at index 'i'
   a = append(a[:i], a[i+1:]...) // ... is a variadic argument
   ```
3. Replace an element with another
   ```
   // Replace at the "i" index with last element
   a[i] = a[len(a)-1]
   ```
4. Get particular elements from a slice
   ```
   // To get elements from 'j' to 'm', we do
   a = a[j:n] // j is included, but n is not
   ```
5. Length & Capacity of Slice
   ```
   len(s)
   cap(s)
   ```
6. Using range
   ```
   s := []int{10, 20, 30, 40}
   for key, value := range s {
       fmt.Println(key, value)
   }
   ```
7. Using for loop
   ```
   for i := 0; i < len(s); i++ {
       fmt.Println(s[i]) // get the value at index "i"
   }
   ```
8. Nested Slice
   ```
   nestedInt := make([][]int, 0) // Slice of a slice of int
   nesteadString := [][]string{slice1, slice2} // slice of slice of string
   ```
#### Maps
Go provides a built-in map type that implements a hash table.

Syntax:
```
map[keyType]valueType
```

Can be defined in 3 ways
```
var sampleMap = map[string]int // using var
sampleMap := map[string]int // shorthand notation
sampleMap := make(map[string]int) // using make 
```


#### Reference
1. https://medium.com/@victorsteven/understanding-data-structures-in-golang-f55205afdcaa
2. https://blog.golang.org/slices-intro


