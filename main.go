package main

import (
	"fmt"

	"github.com/KodeFu/rtc/features"
)

func main() {
	fmt.Println("hello ray!")

	fmt.Println("5 + 3 = ", features.Add(5, 3))
}
