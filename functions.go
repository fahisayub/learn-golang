package main

import ("fmt")

func add(x int, y int) int { 
	// since Go lang is statically typed language you need to specify the type of input and output
	return x + y
}
func main(){
	fmt.Println(add(20,40))
}