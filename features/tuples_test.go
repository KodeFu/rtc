package features

import (
	"testing"

	"sample.com/rtc/utils"
)

func TestTuplePoint(testing *testing.T) {
	var t Tuple
	t.Elements[0] = 4.3
	t.Elements[1] = -4.2
	t.Elements[2] = 3.1
	t.Elements[3] = 1.0

	if t.Elements[0] != 4.3 || t.Elements[1] != -4.2 || t.Elements[2] != 3.1 || t.Elements[3] != 1.0 {
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
	var t Tuple
	t.Elements[0] = 4.3
	t.Elements[1] = -4.2
	t.Elements[2] = 3.1
	t.Elements[3] = 0.0

	if t.Elements[0] != 4.3 || t.Elements[1] != -4.2 || t.Elements[2] != 3.1 || t.Elements[3] != 0.0 {
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
	var t = NewPoint(4, -4, 3)

	if t.Elements[0] != 4 || t.Elements[1] != -4 || t.Elements[2] != 3 || t.Elements[3] != 1.0 {
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
	var t = NewVector(4, -4, 3)

	if t.Elements[0] != 4 || t.Elements[1] != -4 || t.Elements[2] != 3 || t.Elements[3] != 0 {
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
	var t1 = NewTuple(3, -2, 5, 1)
	var t2 = NewTuple(-2, 3, 1, 0)

	r := t1.Add(t2)

	if r.Elements[0] != 1 || r.Elements[1] != 1 || r.Elements[2] != 6 || r.Elements[3] != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestSubTwoPoints(testing *testing.T) {
	var p1 = NewPoint(3, 2, 1)
	var p2 = NewPoint(5, 6, 7)

	r := p1.Sub(p2)

	if r.Elements[0] != -2 || r.Elements[1] != -4 || r.Elements[2] != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubVectorFromPoint(testing *testing.T) {
	var p = NewPoint(3, 2, 1)
	var v = NewVector(5, 6, 7)

	r := p.Sub(v)

	if r.Elements[0] != -2 && r.Elements[1] != -4 && r.Elements[2] != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsPoint(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubTwoVectors(testing *testing.T) {
	var v1 = NewVector(3, 2, 1)
	var v2 = NewVector(5, 6, 7)

	r := v1.Sub(v2)

	if r.Elements[0] != -2 && r.Elements[1] != -4 && r.Elements[2] != -6 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestSubVectorAndZeroVector(testing *testing.T) {
	var z = NewVector(0, 0, 0)
	var v = NewVector(1, -2, 3)

	r := z.Sub(v)

	if r.Elements[0] != -1 && r.Elements[1] != 2 && r.Elements[2] != -3 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestNegateTuple(testing *testing.T) {
	var a = NewTuple(1, -2, 3, -4)

	r := a.Negate()

	if r.Elements[0] != -1 || r.Elements[1] != 2 || r.Elements[2] != -3 || r.Elements[3] != 4 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMultScaler(testing *testing.T) {
	var a = NewTuple(1, -2, 3, -4)

	r := a.Mult(3.5)

	if r.Elements[0] != 3.5 || r.Elements[1] != -7.0 || r.Elements[2] != 10.5 || r.Elements[3] != -14.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMultFraction(testing *testing.T) {
	var a = NewTuple(1, -2, 3, -4)

	r := a.Mult(0.5)

	if r.Elements[0] != 0.5 || r.Elements[1] != -1.0 || r.Elements[2] != 1.5 || r.Elements[3] != -2.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestDiv(testing *testing.T) {
	var a = NewTuple(1, -2, 3, -4)

	r := a.Div(2)

	if r.Elements[0] != 0.5 || r.Elements[1] != -1.0 || r.Elements[2] != 1.5 || r.Elements[3] != -2.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag1(testing *testing.T) {
	var a = NewVector(1, 0, 0)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag2(testing *testing.T) {
	var a = NewVector(0, 1, 0)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag3(testing *testing.T) {
	var a = NewVector(0, 0, 1)

	r := a.Magnitude()

	if r != 1 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestMag4(testing *testing.T) {
	var a = NewVector(1, 2, 3)

	r := a.Magnitude()

	if !utils.Equal(r, 3.7416573867739413) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestNormalize1(testing *testing.T) {
	var a = NewVector(4, 0, 0)

	r := a.Normalize()

	if r.Elements[0] != 1.0 || r.Elements[1] != 0.0 || r.Elements[2] != 0.0 {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestNormalize2(testing *testing.T) {
	var a = NewVector(1, 2, 3)

	r := a.Normalize()

	if !utils.Equal(r.Elements[0], 0.26726) || !utils.Equal(r.Elements[1], 0.53452) || !utils.Equal(r.Elements[2], 0.80178) {
		testing.Errorf("unexpected result %v", r)
	}

	if !IsVector(r) {
		testing.Errorf("not a vector %v", r)
	}
}

func TestDotProduct(testing *testing.T) {
	var a = NewVector(1, 2, 3)
	var b = NewVector(2, 3, 4)

	r := a.Dot(b)

	if r != 20 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestCrossProduct1(testing *testing.T) {
	var a = NewVector(1, 2, 3)
	var b = NewVector(2, 3, 4)

	r := a.Cross(b)

	if r.Elements[0] != -1.0 || r.Elements[1] != 2.0 || r.Elements[2] != -1.0 || !IsVector(r) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestCrossProduct2(testing *testing.T) {
	var a = NewVector(1, 2, 3)
	var b = NewVector(2, 3, 4)

	r := b.Cross(a)

	if r.Elements[0] != 1.0 || r.Elements[1] != -2.0 || r.Elements[2] != 1.0 || !IsVector(r) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestEqual(testing *testing.T) {
	var a = NewTuple(1, 2, 3, 1)
	var b = NewTuple(1, 2, 3, 1)

	if !a.Equals(b) {
		testing.Errorf("unexpected result %v %v", a, b)
	}
}

func TestNotEqual(testing *testing.T) {
	var a = NewTuple(1, 2, 3, 1)
	var b = NewTuple(0, 2, 3, 1)

	if a.Equals(b) {
		testing.Errorf("unexpected result %v %v", a, b)
	}

	b = NewTuple(1, 3, 3, 1)

	if a.Equals(b) {
		testing.Errorf("unexpected result %v %v", a, b)
	}

	b = NewTuple(1, 2, 4, 1)

	if a.Equals(b) {
		testing.Errorf("unexpected result %v %v", a, b)
	}

	b = NewTuple(1, 2, 3, 4)

	if a.Equals(b) {
		testing.Errorf("unexpected result %v %v", a, b)
	}
}
