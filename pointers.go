package main
import ("fmt")

func main(){
	i:=40
	p:=&i // address of i
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p) // this will print the value back i.e value of i
	*p=29// this will update the value in i
	fmt.Println(i)
	

}