package main

import "fmt"

func countmaze(sr , sc , dr , dc int)[]string{
	var hpath []string
	var vpath []string
	if sr == dr && sc == dc {
		temp := []string{""}
		return temp
	}
	if sc < dc{
		hpath = countmaze(sr , sc + 1 , dr , dc)
	}
	if sr < dr {
		vpath = countmaze(sr + 1 , sc ,dr ,dc)
	}
	
	var path [] string
	for _ , h := range hpath{
		path = append(path, "h"+h)
	}
	for _ , v := range vpath{
		path  = append(path, "v"+v)
	}
	return path
}
func main(){
	fmt.Print(countmaze(1 ,1 ,3,3))
}


func minPathSum(grid [][]int) int {
    dp := make([][]int , len(grid))
    for i := 0 ; i < len(dp) ; i++{
        dp[i] = make([]int , len(grid[0]))
    }
    for i := 0 ; i < len(dp) ; i++{
        dp[i][0] = grid[i][0]
    }
     for i := 0 ; i < len(dp[0]) ; i++{
        dp[0][i] = grid[0][i]
    }
    for i := 1 ; i < len(dp) ; i++{
        for j := 1 ; j< len(dp[0]) ; j++{
            dp[i][j] = grid[i][j] + min(dp[i-1][j],dp[i][j-1])
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
    
}
func min(a int , b int) int{
    if a < b {
        return b
    }else{
        return a
    }
}