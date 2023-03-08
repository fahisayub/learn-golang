package main

import ("fmt")

type Person struct{
	name string
	age int
	gender string
}

func main(){
	fmt.Println(Person{"Fahiz",24,"male"})
}