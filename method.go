package main

import ("fmt")

// to make method we need to create a struct first
type rect struct{
	height,width int
}
// we call it as method
func (r rect) area()int{ 
	return r.height*r.width
}

func main(){
	r:=rect{width: 5,height: 10}
	fmt.Println("area:",r.area())

}