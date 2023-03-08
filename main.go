package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	var j = "fahiz"
	var fname, lname = "M", "Fahiz"
	name := "muhammedfahiz" //short decleration
	a, b := 2.2, 3.4        //float32 type
	c, d := 5, 8
	fmt.Println(i)
	fmt.Println(name)
	fmt.Println(fname + lname)
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(a / b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(c % d) // moduls is not possible in float values
	c++
	fmt.Println(c)
	d--
	d+=5
	c-=2
	d/=2
	d%=2
	fmt.Println(d)
	fmt.Println(c==d)
	fmt.Println(c!=d)
	fmt.Println(c>d||c<d)
	fmt.Println(c>=d||c<=d)
	fmt.Println(a>b&&c>d)
	fmt.Println(!(a<b))
	var arr [5]int //defining array with fixed length
	fmt.Println(arr)
	arr[2]=10
	fmt.Println(arr)
	//to define array of unknown length(known as slice)
	var uarr [] int
	fmt.Println(uarr)
	// more efficiently we can define array like
	earr:=[5]int{1,2,3,4,5}
	fmt.Println(earr)

}
