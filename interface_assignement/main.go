package main

import "fmt"

type triangle struct {
	base int
	height int
}
type square struct {
	side int
	aside int
}

type shape interface {
	getArea() float64
}

func main() {
	tr  := triangle{base: 10, height: 10}
	sq := square{side: 10, aside: 10}

	printArea(tr)
	printArea(sq)

}

func (tr triangle)getArea() float64 {
	return 0.5 * float64(tr.base) * float64(tr.height)
}

func (sq square)getArea() float64 {
	return float64(sq.side) * float64(sq.aside)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}