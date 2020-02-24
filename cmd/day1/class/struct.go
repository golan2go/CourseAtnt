package main

import (
	"fmt"
)

type Point struct {
	X, Y float32
}
type A struct {
	val int
}

type B struct {
	obj A
}

func main() {
	structInStruct()
	movePoint()

}

func movePoint() {
	p := Point{20, 40}
	pP(p)
	move(&p)
	pP(p)
	p.move()
	pP(p)

}

func pP(p Point) {
	fmt.Printf("%+v\n", p)
}

func move(p *Point) {
	p.X += 10
	p.Y += 10
}

func (p *Point) move() {
	p.X += 10
	p.Y += 10
}

func structInStruct() {
	a1 := A{1}
	b2 := B{a1}

	pAB(a1, b2)

	a1.val = 33

	pAB(a1, b2)
}

func pAB(a A, b B) {
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

//	fmt.Printf("%v, %+v, %#v\n", val, val, val)
