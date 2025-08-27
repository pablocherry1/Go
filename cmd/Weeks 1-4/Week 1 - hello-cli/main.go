package main

import (
    "fmt"
    "flag"
)

func main() {

    name := flag.String("name", "World", "The name to greet")

    flag.Parse()

    fmt.Println("Hello", *name + "!")

}