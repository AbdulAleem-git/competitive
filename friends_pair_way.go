package main

import "fmt"

func pairingways(n int)int {
	dp := make([] int , n+1)
	dp [1] = 1
	dp [2] = 2
	for i := 3 ; i <= n ; i++{
		dp[i] = dp[i-1] + (i-1) * dp[i-2]
	}
	return dp[n]
}
func main(){
	fmt.Println(pairingways(4))
}