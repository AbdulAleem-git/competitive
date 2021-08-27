package main

import "fmt"
func dfs(i int , n int) {
	if i > n{
		return;
	}

	fmt.Print(i)
	for j := 0 ; j <= 9 ; j++{
		dfs(10*i+j,n)
	}
}
func main(){
	for i:=0 ; i<=9;i++{
		dfs(i,1000)
	}
}
func generate(numRows int) [][]int {
    dp := make([][]int , numRows)
    dp[0] =[]int{1}
    for i := 1 ; i < len(dp) ; i++{
        dp[i] = make([]int , len(dp[i-1])+1)
        for j := 0 ; j < len(dp[i]) ; j++{
            if j == 0 || j == len(dp[i]) - 1 {
                dp[i][j] = 1
            }else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }
    }
    return dp
}