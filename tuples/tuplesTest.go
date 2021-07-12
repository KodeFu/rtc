package tuples

import (
	"fmt"
	"log"
)

func testTuplePoint() {
	var t = Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	if t.X == 4.3 && t.Y == -4.2 && t.Z == 3.1 && t.W == 1.0 {
		if IsPoint(t) {
			log.Println("testTuplePoint passed")
			return
		}
	}

	log.Print("testTuplePoint failed")
}

func testTupleVector() {
	var t = Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	if t.X == 4.3 && t.Y == -4.2 && t.Z == 3.1 && t.W == 1.0 {
		if IsVector(t) {
			log.Println("testTupleVector passed")
			return
		}
	}

	log.Print("testTupleVector failed")
}

func TestStuff() {
	var t = Tuple{X: 1, Y: 1, Z: 1, W: 1}

	if IsPoint(t) {
		fmt.Println("IsPoint")
	}

	if IsVector(t) {
		fmt.Println("IsVector")
	}

	fmt.Println(t)

	t = Tuple{X: 1, Y: 1, Z: 1, W: 0}

	if IsPoint(t) {
		fmt.Println("IsPoint")
	}

	if IsVector(t) {
		fmt.Println("IsVector")
	}

	fmt.Println(t)
}

func TestAll() {
	testTuplePoint()
	testTupleVector()
}
