package main

import (
	"fmt"
)

func targetsum(arr [] int , target , ans int ,subset []int){
	if ans == target{
		fmt.Println(subset)
		return 
	}
	if len(arr) > 1{
		targetsum(arr[1:] , target , ans+arr[0] , append(subset , arr[0]))
		targetsum(arr[1:] , target, ans ,subset)
	}
}
func main(){

	targetsum([]int{10,20,30,40,50},500,0,[]int{})
}