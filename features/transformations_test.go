package features

import (
	"math"
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

func TestRotatingPointAroundXAxis(testing *testing.T) {
	var p = NewPoint(0, 1, 0)
	var half_quarter = Rotation_X(math.Pi / 4)
	var full_quarter = Rotation_X(math.Pi / 2)

	var result_half_quarter = half_quarter.MultiplyByTuple(p)
	var result_full_quarter = full_quarter.MultiplyByTuple(p)

	var expected_half_quarter = NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	var expected_full_quarter = NewPoint(0, 0, 1)

	if !result_half_quarter.Equals(expected_half_quarter) {
		testing.Errorf("unexpected result %v %v", result_half_quarter, expected_half_quarter)
	}

	if !result_full_quarter.Equals(expected_full_quarter) {
		testing.Errorf("unexpected result %v %v", result_full_quarter, expected_full_quarter)
	}
}

func TestRotatingPointAroundXAxisOpposite(testing *testing.T) {
	var p = NewPoint(0, 1, 0)
	var half_quarter = Rotation_X(math.Pi / 4)
	var inv = half_quarter.Inverse()

	var result_half_quarter = inv.MultiplyByTuple(p)
	var expected_half_quarter = NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	if !result_half_quarter.Equals(expected_half_quarter) {
		testing.Errorf("unexpected result %v %v", result_half_quarter, expected_half_quarter)
	}
}

func TestRotatingPointAroundYAxis(testing *testing.T) {
	var p = NewPoint(0, 0, 1)
	var half_quarter = Rotation_Y(math.Pi / 4)
	var full_quarter = Rotation_Y(math.Pi / 2)

	var result_half_quarter = half_quarter.MultiplyByTuple(p)
	var result_full_quarter = full_quarter.MultiplyByTuple(p)

	var expected_half_quarter = NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	var expected_full_quarter = NewPoint(1, 0, 0)

	if !result_half_quarter.Equals(expected_half_quarter) {
		testing.Errorf("unexpected result %v %v", result_half_quarter, expected_half_quarter)
	}

	if !result_full_quarter.Equals(expected_full_quarter) {
		testing.Errorf("unexpected result %v %v", result_full_quarter, expected_full_quarter)
	}
}

func TestRotatingPointAroundZAxis(testing *testing.T) {
	var p = NewPoint(0, 1, 0)
	var half_quarter = Rotation_Z(math.Pi / 4)
	var full_quarter = Rotation_Z(math.Pi / 2)

	var result_half_quarter = half_quarter.MultiplyByTuple(p)
	var result_full_quarter = full_quarter.MultiplyByTuple(p)

	var expected_half_quarter = NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	var expected_full_quarter = NewPoint(-1, 0, 0)

	if !result_half_quarter.Equals(expected_half_quarter) {
		testing.Errorf("unexpected result %v %v", result_half_quarter, expected_half_quarter)
	}

	if !result_full_quarter.Equals(expected_full_quarter) {
		testing.Errorf("unexpected result %v %v", result_full_quarter, expected_full_quarter)
	}
}
