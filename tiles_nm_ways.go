package main

import "fmt"

func tilling_ways(n , m int) int {
	dp := make([]int , n + 1)
	for i := 0 ; i <= n ; i++{
		if i < m {
			dp[i] = 1
		}else if i == m{
			dp[i] = 2
		}else{
			dp[i] = dp[i-1] + dp[i-m]
		}
	}
	return dp[n]
}
func main(){
	fmt.Println(tilling_ways(8,3))
}