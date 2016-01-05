package main

import "fmt"

type animal interface {
	run()
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d *dog) run() {
	fmt.Printf("I am a dog and my name is %s\n", d.name)
}

func (c *cat) run() {
	fmt.Printf("I am a cat and my name is %s\n", c.name)
}

func runwa(a animal) {
	a.run()
}

func main() {
	d := dog{name: "poppy"}
	c := cat{name: "pussycat"}
	runwa(&d)
	runwa(&c)
}
