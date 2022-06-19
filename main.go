package main

import (
	"fmt"
	"os"
)

func main() {
	foo := os.Getenv("APIKEY")
	fmt.Printf("foo: %v\n", foo)
}
