package main

import "fmt"

func maxsum(arr[] int )int{
	include := 0
	exclude := 0
	for i:=0 ; i <len(arr);i++{
		include_new := arr[i] + exclude	
		exclude = max(include , exclude)
		include = include_new
	}
	return max(include , exclude)
}
func max(a int , b int) int {
	if a > b {
		return a 
	}else{
		return b
	}
}

func main(){
	arr := []int{5,10,10,100,5,6}
	fmt.Println(maxsum(arr))
}