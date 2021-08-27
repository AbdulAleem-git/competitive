package main

import "fmt"
func knapsack(cap int ,capicities[]int, profits []int) []int {
	res := make([]int ,cap+1)
	res[0] = 0
	for i:= 1 ; i < len(res);  i++{
		for j := 0 ; j < len(capicities); j++{
			if i >= capicities[j]{
				res[i] = max (res[i] , res[i-capicities[j]]+ profits[j])
 			}
		}
	}
	return res
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