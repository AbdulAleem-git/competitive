package main

import (
	"fmt"
	"math"
)
var max int = int(math.Pow(2, -31))

func print(arr []int , idx int){
	if idx > len(arr)-1{
		return
	}
	print(arr, idx+1)
	fmt.Println(arr[idx])

}
func maxvalue(arr []int , idx int ) int{	
	if idx > len(arr) -1 {
		return max
	}
	if max < arr[idx]{
		max = arr[idx]
	}
	return maxvalue(arr, idx+1)
}

func main(){
	// print([]int{1,2,3,4,5,6,7,8,9,10} , 0)
	fmt.Println(maxvalue([]int{1,2,3,4,5,6,17,8,9,10} , 0))
	fmt.Println(maxvalue([]int{} , 0))
}