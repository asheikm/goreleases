package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["1"] = "one"
	m["2"] = "two"
	m["3"] = "three"
	m["4"] = "four"
	fmt.Println("*********** clear with maps **********************")
	fmt.Println("Printing values of map m before making call to clear(m)")
	for i := range m {
		fmt.Println(m[i])
	}
	fmt.Println("Length of values of map m before making call to clear(m): ", len(m))
	clear(m)
	fmt.Println("Length of values of map m after making call to clear(m): ", len(m))
	fmt.Println("Printing values of map m after making call to clear(m)")
	for i := range m {
		fmt.Println(m[i])
	}
	fmt.Println("*********** Clear with slices **********************")
	s := make([]int, 3, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("Printing values of slice s before making call to clear(s)")

	fmt.Println("Length and capacity of values of slice s before making call to clear(s): ", len(s), cap(s))
	for i := range s {
		fmt.Println(s[i])
	}
	clear(s)
	fmt.Println("Length and capacity of values of slice s before making call to clear(s): ", len(s), cap(s))
	fmt.Println("Printing values of slice s before making call to clear(s)")
	for i := range s {
		fmt.Println(s[i])
	}
	fmt.Println("*********** Clear with types with slices **********************")
	type Point struct {
		X int
		Y int
	}
	pointSlice := []Point{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
	}

	fmt.Println("Length and capacity of values of type struct pointSlice before making call to clear(pointSlice): ", len(pointSlice), cap(pointSlice))
	fmt.Println("Printing values of type slice pointSlice before making call to clear(pointSlice)")
	for i := range pointSlice {
		fmt.Println(pointSlice[i])
	}
	clear(pointSlice)
	fmt.Println("Length and capacity of values of type struct pointSlice before making call to clear(pointSlice): ", len(pointSlice), cap(pointSlice))
	fmt.Println("Printing values of type slice pointSlice after making call to clear(pointSlice)")
	for i := range pointSlice {
		fmt.Println(pointSlice[i])
	}
	fmt.Println("*********** Clear with types with map **********************")
	type MyMap map[string]int
	// Initialize an instance of MyMap
	myMapInstance := MyMap{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println("Length values of type map MyMap before making call to clear(myMapInstance): ", len(myMapInstance))
	fmt.Println("Printing values of type map myMapInstance before making call to clear(myMapInstance)")
	for i := range myMapInstance {
		fmt.Println(myMapInstance[i])
	}
	clear(myMapInstance)
	fmt.Println("Length values of type map MyMap before making call to clear(myMapInstance): ", len(myMapInstance))
	fmt.Println("Printing values of type map myMapInstance after making call to clear(myMapInstance)")
	for i := range myMapInstance {
		fmt.Println(myMapInstance[i])
	}
}
/*************************************************
*********** clear with maps **********************
Printing values of map m before making call to clear(m)
one
two
three
four
Length of values of map m before making call to clear(m):  4
Length of values of map m after making call to clear(m):  0
Printing values of map m after making call to clear(m)
*********** Clear with slices **********************
Printing values of slice s before making call to clear(s)
Length and capacity of values of slice s before making call to clear(s):  3 3
1
2
3
Length and capacity of values of slice s before making call to clear(s):  3 3
Printing values of slice s before making call to clear(s)
0
0
0
*********** Clear with types with slices **********************
Length and capacity of values of type struct pointSlice before making call to clear(pointSlice):  2 2
Length and capacity of values of type struct pointSlice before making call to clear(pointSlice):  2 2
*********** Clear with types with map **********************
Length values of type map MyMap before making call to clear(myMapInstance):  2
Length values of type map MyMap before making call to clear(myMapInstance):  0

Program exited.*/
