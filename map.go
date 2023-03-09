package main
import ("fmt")

func main(){
	m:=make(map[string]int)
	m["a"]=3
	m["b"]=2
	fmt.Println(m)
	delete(m,"b")
	fmt.Println(m)
	nm:=map[string]int{"a":10, "b":12}
	fmt.Println(len(nm))
	fmt.Println(nm)
}