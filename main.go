package main

import (
	"fmt"

	"sample.com/rtc/features"
)

type Projectile struct {
	pPosition features.Tuple
	vVelocity features.Tuple
}

type Environment struct {
	vGravity features.Tuple
	vWind    features.Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	var p Projectile
	p.pPosition = proj.pPosition.Add(proj.vVelocity)
	p.vVelocity = proj.vVelocity.Add(env.vGravity).Add(env.vWind)
	return p
}

func main() {
	fmt.Println("hello ray!")
	/*
		a := features.Tuple{X: 3, Y: 4, Z: 2, W: 1}
		b := features.Tuple{X: 3, Y: 4, Z: 2, W: 0}
		c := a.Add(b)

		fmt.Println(c)

		fmt.Println(a.Add(b))

		p := Projectile{features.Point(0, 1, 0), features.Vector(1, 1, 0).Normalize()}
		e := Environment{features.Vector(0, -0.1, 0), features.Vector(-0.01, 0, 0)}

		fmt.Println("p", p)
		fmt.Println("e", e)

		newP := Tick(e, p)
		fmt.Println("newP", newP)
		newP = Tick(e, newP)
		fmt.Println("newP", newP)
		newP = Tick(e, newP)
		fmt.Println("newP", newP)
		newP = Tick(e, newP)
		fmt.Println("newP", newP)
	*/
	var canvas = features.NewCanvas(5, 3)
	//fmt.Println(canvas.Pixels)
	canvas.CanvasToPPM()
	/*
		var matrix = features.NewMatrix(2, 4)
		fmt.Println(matrix)

		matrix = features.NewMatrix(2, 3)
		fmt.Println(matrix)

		matrix = features.NewMatrix(9, 3)
		fmt.Println(matrix)
	*/

	var m = features.NewMatrix(0, 0)
	fmt.Printf("size %d, size %d", len(m.Elements), len(m.Elements[0]))
	fmt.Printf("size %d, size %d", m.NumRows(), m.NumCols())
}
