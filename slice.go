package main
 import ("fmt")

 func main(){
 // we can define a slice using make fucntion
 arr:=make([]string,2)
 fmt.Println(arr)
 arr[0]="a"
 arr[1]="b"
 fmt.Println(arr)
 arr = append(arr, "x")
 fmt.Println(arr)
 fmt.Println(len(arr))


nums:=[]int{1,2,3,4}
for i,num:=range nums { // we can access both index and value using range
	fmt.Println(i,num)
}

 }