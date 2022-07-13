package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println("Prefix", strings.HasPrefix("hello world", "hello"))

	t := time.Now()
	fmt.Println("t", t.Day())
}
