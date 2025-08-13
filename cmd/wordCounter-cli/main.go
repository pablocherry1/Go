package main

import (
	"fmt"
	"errors"
)

func main() {
	
	myError := errors.New("Something went wrong in the funtion")
	fmt.Println(myError)
}
