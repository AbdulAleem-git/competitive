package main

import (
	"fmt"
	"math"
)
func mincost(arr [][]int) int{
	dp := make ([][]int, len(arr))
	for i:=0 ; i < len(dp); i++{
		dp[i] = make([]int, len(arr[0]))
	}
	least := int(math.Pow(2,32) - 1 )
	sleast := int(math.Pow(2,32) - 1)
	for j := 0 ; j < len(arr[0]); j++	{
		dp[0][j] = arr[0][j]
		if least >= arr[0][j] {
			sleast= least
			least = arr[0][j]
		}else if sleast >= arr[0][j]{
			sleast = arr[0][j]
		}
	}
	for i := 1 ; i < len(dp) ; i++{
		nleast := int(math.Pow(2,32) - 1 )
		nsleast := int(math.Pow(2,32) - 1)
		for j := 0 ; j < len(dp[0]); j++{
			if dp[i-1][j] == least{
				dp[i][j] = arr[i][j] + sleast
			}else{
				dp[i][j] = arr[i][j] + least
			}
			if nleast >= dp[i][j] {
				nsleast= nleast
				nleast = dp[i][j]
			}else if nsleast >= dp[i][j]{
				nsleast = dp[i][j]
			}
		}
		least = nleast
		sleast = nsleast
	}
return least
}
func main(){
	arr := [][]int{
		{1,5,7,2,3,4},
		{5,8,4,3,6,1},
		{3,2,9,7,2,3},
		{1,2,4,9,1,7},
	}
	fmt.Println(mincost(arr))
}