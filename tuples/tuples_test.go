package tuples

import (
	"testing"
)

func TestTuplePoint(testing *testing.T) {
	var t = Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	if t.X != 4.3 || t.Y != -4.2 || t.Z != 3.1 || t.W != 1.0 {
		testing.Errorf("unexpected result %v", t)
	}
	if !IsPoint(t) {
		testing.Errorf("not a point %v", t)
	}
	if IsVector(t) {
		testing.Errorf("is vector, point expected %v", t)
	}
}

func TestTupleVector(testing *testing.T) {
	var t = Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	if t.X != 4.3 || t.Y != -4.2 || t.Z != 3.1 || t.W != 0.0 {
		testing.Errorf("unexpected result %v", t)
	}
	if !IsVector(t) {
		testing.Errorf("not a vector %v", t)
	}
	if IsPoint(t) {
		testing.Errorf("is point, vector expected %v", t)
	}
}

func TestIsPoint(testing *testing.T) {
	var t = Point(4.3, -4.2, 3.1)

	if t.X != 4.3 || t.Y != -4.2 || t.Z != 3.1 || t.W != 1.0 {
		testing.Errorf("unexpected result %v", t)
	}
	if !IsPoint(t) {
		testing.Errorf("not a point %v", t)
	}
	if IsVector(t) {
		testing.Errorf("is vector, point expected %v", t)
	}
}

func TestStuff(testing *testing.T) {
	var t = Tuple{X: 1, Y: 1, Z: 1, W: 1}

	if IsPoint(t) {
	}

	if IsVector(t) {
	}

	t = Tuple{X: 1, Y: 1, Z: 1, W: 0}

	if IsPoint(t) {

	}

	if IsVector(t) {
	}

}
