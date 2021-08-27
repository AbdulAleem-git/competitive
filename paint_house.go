package main

import (
	"fmt"
	"math"
)
func mincost(arr [][]int) int{
	
	red := arr[0]
	blue := arr[1]
	green := arr[2]
	redc := red[0]
	bluec := blue[0]
	greenc := green[0]
	for i := 1; i < len(arr[0]) ; i++{
		redc_new := min(bluec , greenc) + red[i]
		bluec_new := min(redc , greenc) + blue[i]
		greenc_new := min(bluec , redc) + green[i]
		
		redc = redc_new
		bluec = bluec_new
		greenc = greenc_new

	}
	return min(redc ,bluec ,greenc)
}
func min(a ...int) int{
	min := int(math.Pow(2 , 32)) - 1
	for _, value := range a{
		if min > value{
			min = value
		}
	}
	return min
}

func main(){
	arr  := [][]int{
		{1,5,3,1},
		{5,8,2,2},
		{7,4,9,4},
	}
	fmt.Println(mincost(arr))
}