package features

import (
	"testing"
)

func TestTranslationMultiply(testing *testing.T) {
	var transform = Translation(5, -3, 2)
	var p = NewPoint(-3, 4, 5)

	var result = transform.MultiplyByTuple(p)

	var expected = NewPoint(2, 1, 7)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestTranslationMultiplyInverse(testing *testing.T) {
	var transform = Translation(5, -3, 2)
	var inv = transform.Inverse()
	var p = NewPoint(-3, 4, 5)

	var result = inv.MultiplyByTuple(p)

	var expected = NewPoint(-8, 7, 3)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestTranslationVectors(testing *testing.T) {
	var transform = Translation(5, -3, 2)
	var v = NewVector(-3, 4, 5)

	var result = transform.MultiplyByTuple(v)

	if !result.Equals(v) {
		testing.Errorf("unexpected result %v %v", result, v)
	}
}

func TestScalingPoint(testing *testing.T) {
	var transform = Scaling(2, 3, 4)
	var p = NewPoint(-4, 6, 8)

	var result = transform.MultiplyByTuple(p)

	var expected = NewPoint(-8, 18, 32)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestScalingVector(testing *testing.T) {
	var transform = Scaling(2, 3, 4)
	var v = NewVector(-4, 6, 8)

	var result = transform.MultiplyByTuple(v)

	var expected = NewVector(-8, 18, 32)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}

func TestScalingMultiplyInverse(testing *testing.T) {
	var transform = Scaling(2, 3, 4)
	var inv = transform.Inverse()
	var v = NewVector(-4, 6, 8)

	var result = inv.MultiplyByTuple(v)

	var expected = NewVector(-2, 2, 2)

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", result, expected)
	}
}
