package main

import "fmt"

func main() {
	// int
	var x, y int
	x, y = 5, 10
	a := min(x)
	b := max(x)
	c := min(x, y)
	d := max(x, y)
	e := min(a, b, c, d)
	f := max(a, b, c, d)
	fmt.Println(a, b, c, d, e, f)

	// string
	g := "something"
	h := "nothing"
	i := "something is better than nothing"
	j := min(g, h, i)
	k := max(g, h, i)
	fmt.Println(j, k)

}
