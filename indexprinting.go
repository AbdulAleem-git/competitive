package main

import "fmt"
var res []int
func printindex(arr [] int , x , idx int)( []int){
	if idx > len(arr)-1{
		return res
	}
	if arr[idx] == x {
		res = append(res, idx)
	}
	return printindex(arr , x , idx+1)	
}
func main(){
	fmt.Println(printindex([]int{2,5,6,7,6,6,2,5,6,6,4,3,6} , 6 , 0))
}