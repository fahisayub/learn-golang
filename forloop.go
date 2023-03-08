package main
 import ("fmt")

 func main() {

	sum:=0
	for i:=0;i<10;i++{
		sum+=i;
	}
	fmt.Println(sum)
// in go there is no concept like while instead we can do like this:
for sum>10 {
	sum++;
}
for{
	sum++ // this is infinite loop
	break;// to stop it
}

 }