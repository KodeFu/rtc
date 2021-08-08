package tuples

import (
	"testing"

	"sample.com/rtc/utils"
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
	var t = Point(4, -4, 3)

	if t.X != 4 || t.Y != -4 || t.Z != 3 || t.W != 1.0 {
		testing.Errorf("unexpected result %v", t)
	}
	if !IsPoint(t) {
		testing.Errorf("not a point %v", t)
	}
	if IsVector(t) {
		testing.Errorf("is vector, point expected %v", t)
	}
}

func TestIsVector(testing *testing.T) {
	var t = Vector(4, -4, 3)

	if t.X != 4 || t.Y != -4 || t.Z != 3 || t.W != 0 {
		testing.Errorf("unexpected result %v", t)
	}
	if !IsVector(t) {
		testing.Errorf("not a vector %v", t)
	}
	if IsPoint(t) {
		testing.Errorf("is point, vector expected %v", t)
	}
}

func TestAdd(testing *testing.T) {
	var t1 = Tuple{3, -2, 5, 1}
	var t2 = Tuple{-2, 3, 1, 0}

	r := t1.Add(t2)

	if r.X != 1 || r.Y != 1 || r.Z != 6 || r.W != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestSubTwoPoints(testing *testing.T) {
	var p1 = Point(3, 2, 1)
	var p2 = Point(5, 6, 7)

	r := p1.Sub(p2)

	if r.X != -2 || r.Y != -4 || r.Z != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubVectorFromPoint(testing *testing.T) {
	var p = Point(3, 2, 1)
	var v = Vector(5, 6, 7)

	r := p.Sub(v)

	if r.X != -2 && r.Y != -4 && r.Z != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsPoint(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubTwoVectors(testing *testing.T) {
	var v1 = Vector(3, 2, 1)
	var v2 = Vector(5, 6, 7)

	r := v1.Sub(v2)

	if r.X != -2 && r.Y != -4 && r.Z != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubVectorAndZeroVector(testing *testing.T) {
	var z = Vector(0, 0, 0)
	var v = Vector(1, -2, 3)

	r := z.Sub(v)

	if r.X != -1 && r.Y != 2 && r.Z != -3 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestNegateTuple(testing *testing.T) {
	var a = Tuple{1, -2, 3, -4}

	r := a.Negate()

	if r.X != -1 || r.Y != 2 || r.Z != -3 || r.W != 4 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMultScaler(testing *testing.T) {
	var a = Tuple{1, -2, 3, -4}

	r := a.Mult(3.5)

	if r.X != 3.5 || r.Y != -7.0 || r.Z != 10.5 || r.W != -14.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMultFraction(testing *testing.T) {
	var a = Tuple{1, -2, 3, -4}

	r := a.Mult(0.5)

	if r.X != 0.5 || r.Y != -1.0 || r.Z != 1.5 || r.W != -2.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestDiv(testing *testing.T) {
	var a = Tuple{1, -2, 3, -4}

	r := a.Div(2)

	if r.X != 0.5 || r.Y != -1.0 || r.Z != 1.5 || r.W != -2.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag1(testing *testing.T) {
	var a = Vector(1, 0, 0)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag2(testing *testing.T) {
	var a = Vector(0, 1, 0)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag3(testing *testing.T) {
	var a = Vector(0, 0, 1)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag4(testing *testing.T) {
	var a = Vector(1, 2, 3)

	r := a.Magnitude()

	if !utils.Equal(r, 3.7416573867739413) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestNormalize1(testing *testing.T) {
	var a = Vector(4, 0, 0)

	r := a.Normalize()

	if r.X != 1.0 || r.Y != 0.0 || r.Z != 0.0 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestNormalize2(testing *testing.T) {
	var a = Vector(1, 2, 3)

	r := a.Normalize()

	if !utils.Equal(r.X, 0.26726) || !utils.Equal(r.Y, 0.53452) || !utils.Equal(r.Z, 0.80178) {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestDotProduct(testing *testing.T) {
	var a = Vector(1, 2, 3)
	var b = Vector(2, 3, 4)

	r := a.Dot(b)

	if r != 20 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestCrossProduct1(testing *testing.T) {
	var a = Vector(1, 2, 3)
	var b = Vector(2, 3, 4)

	r := a.Cross(b)

	if r.X != -1.0 || r.Y != 2.0 || r.Z != -1.0 || !IsVector(r) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestCrossProduct2(testing *testing.T) {
	var a = Vector(1, 2, 3)
	var b = Vector(2, 3, 4)

	r := b.Cross(a)

	if r.X != 1.0 || r.Y != -2.0 || r.Z != 1.0 || !IsVector(r) {
		testing.Errorf("unexpected result %v", r)
	}
}
