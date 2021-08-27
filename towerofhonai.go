package main

import "fmt"
var k int
func toh(n int , tid1 , tid2 , tid3 string){
	if n == 0 {
		return
	}
	k++
	toh(n-1 ,tid1 , tid3 , tid2 )
	fmt.Printf("%d[%s -> %s]\n",n,tid1,tid2)
	toh(n-1,tid3,tid2 ,tid1 )
}
func main(){
	toh(3, "A","B","C")
	fmt.Println(k)
}