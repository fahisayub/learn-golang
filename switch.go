package main

import ("fmt")

func main(){
	fmt.Println("Start")
	defer fmt.Println("Done") // this only runs at the very end because of defer keyword in go lang
	val:=1
	switch val {
	case 1:
		fmt.Println("value is One")
	case 2:
		fmt.Println("value is Two")
	default:
		fmt.Println("value is Unknown")
	}
 
}