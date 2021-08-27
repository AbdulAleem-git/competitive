package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func maxProfit(prices []int , k int) int{
	dp := make([][]int , k+1)
	for i , _ := range dp {
		dp[i] = make([]int , len(prices))
	}
	for i := 1 ; i < len(dp) ; i++{
		for j := 1 ; j < len(dp[0]) ; j++{
			m := dp[i][j-1]
			for k := j-1 ; k >0 ; k-- {
				m = max(m , dp[i-1][j-1]+prices[j] - prices[k])
			}
			dp[i][j] = m
		}
	}
	fmt.Println(dp)
	return dp[len(dp) - 1][len(dp[0]) - 1]
}
func max(a int , b int )int {
	if a > b {
		return a
	}else {
		return b
	}
}
func main(){
	prices := []int{9,6,7,6,3,8}
	fmt.Printf("%d\n" ,maxProfit(prices, 3))
	var arraysize int 
	fmt.Scanf("%d",&arraysize)
	var prices2 string
	
	reader := bufio.NewReader(os.Stdin)
	prices2 ,_  = reader.ReadString('\n')
	finalprices := strings.Split(strings.Trim(prices2, " ") , " ")
	fmt.Println(finalprices)
	fmt.Println(len(finalprices))
	var pricesv [] int
	for _, v := range finalprices{
		v = strings.Trim(v , " ")
		i ,_ := strconv.Atoi(v)
		pricesv = append(pricesv, i)
	}
	fmt.Println(pricesv)
}