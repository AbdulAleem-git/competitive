package main

import "fmt"

func checksubsetsum(arr[]int , n int) bool{
	s := len (arr)
	res := make([][]bool , s+1)
	for  i := 0 ; i < len(res) ; i++{
		res[i] = make([]bool , n+1)
	}
	for i:= 0 ; i < len(res) ; i++{
		for j:=0; j < len(res[0]); j++{
			if i == 0 && j == 0 {
				res[i][j] = true
			}else if i == 0 {
				res[i][j] = false
			}else if j == 0 {
				res[i][j] = true
			}else if arr[i-1] > j{
				res[i][j] = res[i-1][j]
			}else{
				res[i][j] = res[i-1][j] || res[i-1][j-arr[i-1]]
			}
		}
	}
	return res[len(res)-1][len(res[0])-1]
}
func main(){
	arr := []int{4,2,7,5,2}
	n := 1
	fmt.Println(checksubsetsum(arr , n))
}