package features

type Color struct {
	red, green, blue float64
}

func NewColor(r, g, b float64) Color {
	var c Color
	c.red = r
	c.green = g
	c.blue = b

	return c
}

func (a Color) Add(b Color) Color {
	var r Color
	r.red = a.red + b.red
	r.green = a.green + b.green
	r.blue = a.blue + b.blue

	return r
}

func (a Color) Sub(b Color) Color {
	var r Color
	r.red = a.red - b.red
	r.green = a.green - b.green
	r.blue = a.blue - b.blue

	return r
}

func (a Color) Mult(b float64) Color {
	var r Color
	r.red = a.red * b
	r.green = a.green * b
	r.blue = a.blue * b

	return r
}

func (a Color) HadamardProduct(b Color) Color {
	var r Color
	r.red = a.red * b.red
	r.green = a.green * b.green
	r.blue = a.blue * b.blue

	return r
}
