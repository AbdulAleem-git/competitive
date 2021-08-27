package main

import (
	"fmt"
	"math"
)

func maxgold(arr [][] int ) int {
	n := len(arr)
	m := len(arr[0])
	res := make([][]int , n)
	for i := 0 ; i < n ; i++ {
		res[i] = make([]int, m)
	}
	
	for j :=m-1 ; j >= 0 ; j--{
		for i:= n-1 ; i >=0 ; i--{
			if j == m-1{
				res[i][j] = arr[i][j]
			}else if i == 0{
				res[i][j] = max(res[i][j+1],res[i+1][j+1]) + arr[i][j]
			}else if i == n-1{
				res[i][j] = max(res[i-1][j+1],res[i][j+1]) + arr[i][j]
			}else {
				res[i][j] = max(res[i-1][j+1],res[i][j+1],res[i+1][j+1]) + arr[i][j]
			}
		}
	}
	finalresult := make([]int , n)
	for i:=0 ; i < n ; i++{
		finalresult = append(finalresult, res[i][0])
	}
	max := finalresult[0]
	for _, value := range finalresult{
		if max < value{
			max = value
		}
	}
	return max
}
func max(nums ...int)int{
	max := int(math.Pow(2,-32))
	for _, value := range nums{
		if max < value{
			max = value
		}
	}
	return max
}
func main(){
	arr := [][]int{
		{0,1,4,2,8,2},
		{4,3,6,5,0,4},
		{1,2,4,1,4,6},
		{2,0,7,3,2,2},
		{3,1,5,9,2,4},
		{2,7,0,8,5,1},
	}
	fmt.Println(maxgold(arr))
}