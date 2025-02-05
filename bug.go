```go
package main

import "fmt"

func main() {

    // This map is uninitialized.
    var myMap map[string]int

    // Trying to access a key in an uninitialized map will cause a panic.
    fmt.Println(myMap["key"])
}