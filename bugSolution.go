```go
package main

import "fmt"

func main() {

    // Initialize the map using make
    myMap := make(map[string]int)

    //Or Initialize the map using map literal
    //myMap := map[string]int{}

    // Now it's safe to access keys
    myMap["key"] = 10
    fmt.Println(myMap["key"])
}```