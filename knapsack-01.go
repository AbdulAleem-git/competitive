package main

import "fmt"
func knapsack(cap int ,capicities[]int, profits []int) int {
	res := make([][]int ,len(profits)+1)
	for i,_ := range res{
		res[i] = make([]int , cap+1)
	}
	for i:= 1 ; i < len(res);  i++{
		for j := 1 ; j < len(res[0]); j++{
			if j >= capicities[i-1]{
				res[i][j] = max (res[i-1][j], res[i-1][j-capicities[i-1]]+profits[i-1])
 			}else{
				res[i][j] = res[i-1][j]
			}
		}
	}
	return res[len(res)-1][len(res[0])-1]
}
func max( a int , b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}
func main(){
	capicities := []int{2,5,1,3,4}
	profits := [] int{15,14,10,45,30}
	cap := 7
	fmt.Println(knapsack(cap,capicities,profits))
}