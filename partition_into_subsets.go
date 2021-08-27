package main

import "fmt"
func createteams(n , k int) int{
	dp := make ([][]int , k+1)
	for i := 0 ; i < len(dp) ; i++{
		dp[i] = make([]int , n+1)
	}
	for i := 0 ; i < len(dp) ; i++{
		for j:=0 ; j < len(dp[0]) ; j++{
			if i == 0 || j==0 || i > j {
				dp[i][j] = 0 
			}else if i == j{
				dp[i][j] = 1
			}else {
				dp[i][j] = i * dp[i][j-1] + dp[i-1][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
func main(){
	fmt.Println(createteams(6,5))
}