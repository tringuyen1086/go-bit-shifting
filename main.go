package main

import "fmt"

const (
	// iota used to list elements in an increment order
	_ = iota // blank _ == 0
	a        // a == 1
	b        // b == 2
	c        // c == 3
	d        // d == 4
	e        // e == 5
	f        // f == 6
	g        // g == 7
)

func main() {
	fmt.Println(a, b, c, d, e, f)

	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)

	fmt.Printf("%d \t %b\n", 1<<7, 1<<7)

	fmt.Printf("%d \t %b\n", 64, 64)
	fmt.Printf("%d \t %b\n", 64>>1, 64>>1)
	fmt.Printf("%d \t %b\n", 64>>2, 64>>2)
	fmt.Printf("%d \t %b\n", 64>>3, 64>>3)
	fmt.Printf("%d \t %b\n", 64>>4, 64>>4)
	fmt.Printf("%d \t %b\n", 64>>5, 64>>5)
	fmt.Printf("%d \t %b\n", 64>>6, 64>>6)

	fmt.Println("Iota Display")

	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

	fmt.Printf("%d \t %b\n", 1<<g, 1<<g)

	fmt.Printf("%d \t %b\n", 64, 64)
	fmt.Printf("%d \t %b\n", 64>>a, 64>>a)
	fmt.Printf("%d \t %b\n", 64>>b, 64>>b)
	fmt.Printf("%d \t %b\n", 64>>c, 64>>c)
	fmt.Printf("%d \t %b\n", 64>>d, 64>>d)
	fmt.Printf("%d \t %b\n", 64>>e, 64>>e)
	fmt.Printf("%d \t %b\n", 64>>f, 64>>f)
}
