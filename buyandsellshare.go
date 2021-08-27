package main

import (
	"fmt"
	"math"
)

func calculatemaxprofit(arr [] int ) int {
	lp := int(math.Pow(2,32) - 1)
	tp := 0 
	for i := 0 ; i  < len(arr) ; i++{
		if lp > arr[i]{
			lp = arr[i]
		}
		cp := arr[i] - lp
		if cp > tp{
			tp = cp
		}
	}
	return tp
}
func main(){
	arr := []int{11,6,7,19,4,1,6,18,4,23}
	fmt.Println(calculatemaxprofit(arr))
}