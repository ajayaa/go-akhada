package main

import "fmt"

type point struct {
	x int
	y int
}

func distance_to_origin(p *point) int {
	//return math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return p.x*p.x + p.y*p.y
}

func main() {
	p1 := point{x: 19}
	p2 := point{x: 10, y: 10}
	fmt.Println("p1 is: ", p1)
	fmt.Println("&p2 is: ", &p2)
	fmt.Println("x in p2 is ", (&p2).x)
	fmt.Println("Distance of origin to p2 is: ", distance_to_origin(&p2))
}
