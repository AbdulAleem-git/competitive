package main

import "fmt"

func numberofways(n int)int{
	if n == 0 {
		return 0 
	}else if n == 1{
		return 1
	}else if n == 2 {
		return 2
	}else{
		dp := make([]int, n+1)
		dp[1] = 1
		dp[2] = 2
		for i := 3 ; i < len(dp); i++{
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	}
}
func main(){
	fmt.Println(numberofways(4))
}