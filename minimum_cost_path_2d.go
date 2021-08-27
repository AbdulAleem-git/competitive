package main

import "fmt"

func mincostpath(arr [][]int)int{
	n := len(arr)
	m := len(arr[0])
	//Note-1: The 2D Array Initialization
	res := make([][]int , n)
	for i := 0 ; i < n ; i++{
		res[i] = make([]int ,m)
	}
	// till here

	//implementing using dp
	for i:=n-1 ; i >= 0 ; i--{
		for j:=n-1; j>=0; j--{
			if i == n-1 && j == n-1{
				res[n-1][m-1] = arr[n-1][m-1]
			}else if i == n-1 {
				res[i][j] = res[i][j+1] + arr[i][j]
			}else if j == n-1 {
				res[i][j] = res[i+1][j] + arr[i][j]
			}else{
				res[i][j] = min(res[i+1][j],res[i][j+1]) + arr[i][j]
			}			
		}
	}
	return res[0][0]
}
func min(a int , b int)int{
	if a < b{
		return a
	}else {
		return b
	}
}
func main(){
	arr := [][]int{
		{2,8,4,1,6,4,2},
		{6,0,9,5,3,8,5},
		{1,4,3,4,0,6,5},
		{6,4,7,2,4,6,1},
		{1,0,3,7,1,2,7},
		{1,5,3,2,3,0,9},
		{2,2,5,1,9,8,2},
		{2,2,5,1,9,8,2},
	}
	fmt.Println(mincostpath(arr))
}
