package main

import (
	"fmt"
	"math"
)
// since area and perimeter are very different for circle and rectangle,
//so to handle it and make it sime we will create an abstraction layer
//this helps to create libraries
type geometry interface{
	area() float64
	perimeter() float64
}
// using geometry by making a fucntion
// we can pass any geometry circle or rectangle and can add more without thinking about the whats happening inside
func measure(g geometry){
	fmt.Println("area:",g.area())
	fmt.Println("perimeter:",g.perimeter())
}

 type rect struct {
	width,height float64
 }
 func (r rect) area()float64{
	return r.height*r.width
 }
 func (r rect) perimeter()float64{
	return 2*(r.height+r.width)
 }
 type circle struct {
	radius float64
 }
 func (c circle)area()float64{
	return math.Pi*c.radius*c.radius
 }
 func (c circle) perimeter()float64{
	return 2*math.Pi*c.radius
 }

func main(){
	r:=rect{height: 20,width: 10}
	c:=circle{radius:30}
	measure(r)
	measure(c)

}