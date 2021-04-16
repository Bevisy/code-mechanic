package main

import "fmt"

type cat interface {
	eat(food string) string
	drink(water string) string
	meow()
}

type kongCat struct {
}

func (p *kongCat) eat(food string) string {
	return "eat " + food
}

func (p *kongCat) drink(water string) string {
	return "drink " + water
}

func (p *kongCat) meow() {
	fmt.Println("meow")
}

func main() {
	c := kongCat{}
	c.meow()
	fmt.Printf(c.drink("hot water"))
	fmt.Sprintf(c.eat("cat food"))
}
