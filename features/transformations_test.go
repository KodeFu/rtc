package features

import (
	"testing"
)

func TestTranslationMultiply(testing *testing.T) {
	var transform = Translation(5.0, -3.0, 2.0)
	var p = NewPoint(-3.0, 4.0, 5.0)

	var result = transform.MultiplyByTuple(p)

	var expected = NewPoint(2, 1, 7)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestTranslationMultiplyInverse(testing *testing.T) {
	var transform = Translation(5.0, -3.0, 2.0)
	var inv = transform.Inverse()
	var p = NewPoint(-3.0, 4.0, 5.0)

	var result = inv.MultiplyByTuple(p)

	var expected = NewPoint(-8, 7, 3)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestTranslationVectors(testing *testing.T) {
	var transform = Translation(5.0, -3.0, 2.0)
	var v = NewVector(-3.0, 4.0, 5.0)

	var result = transform.MultiplyByTuple(v)

	if !result.Equals(v) {
		testing.Errorf("unexpected result %v %v", result, v)
	}
}
