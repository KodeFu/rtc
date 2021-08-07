package main

import (
	"fmt"

	"sample.com/rtc/tuples"
)

func main() {
	fmt.Println("hello ray!")

	a := tuples.Tuple{X: 3, Y: 4, Z: 2, W: 1}
	b := tuples.Tuple{X: 3, Y: 4, Z: 2, W: 0}
	c := a.Add(b)

	fmt.Println(c)

	fmt.Println(a.Add(b))
}
