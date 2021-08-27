package main

import (
	"fmt"
	"strconv"
)

func numberofencoding(str string) int {
		dp := make([]int , len(str))
		dp [0] = 1
		for i := 1 ; i < len(dp) ; i++{
			if str[i] == '0' && str[i-1] == '0'{
				dp[i] = 0
			}else if str[i] == '0' && str[i-1] != '0'{
				if str[i-1] == '1' || str[i-1] =='2' {
					if i >=2 {
						dp[i] = dp[i-2]
					}else{
						dp[i] = 1
					}
				}else{
					dp[i] = 0
				}
			}else if str[i] != '0' && str[i-1] == '0'{
				dp[i] = dp [i-1]	
			}else{
				if value , _ := strconv.Atoi(str[i-1:i+1]); value <= 26{
					if i >= 2 {
						dp[i] = dp[i-1] + dp[i-2]
					}else{
						dp[i] = dp[i-1] +1
					}
				}else{
					dp[i] = dp[i-1]
				}
			}
		}
		return dp[len(dp)-1]
}
func main(){
	fmt.Println(numberofencoding("21223"))
	fmt.Println(numberofencoding("21120"))
	fmt.Println(numberofencoding("210"))
}