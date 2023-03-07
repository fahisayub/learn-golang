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
	sum:=0
	for i:=0;i<10;i++{
		sum+=i;
	}
	fmt.Println(sum)
// in go there is no concept like 
}
