package main

import "fmt"

func coins_change_comb(arr [] int , n int) int {
	res := make([]int , n+1)
	res[0] =  1
	for i := 0 ; i < len(arr) ; i++ {
		for j := 1 ; j < n+1 ; j++ {
			if j >= arr[i]{
				res[j] = res[j - arr[i]] + res[j]
			}
		}

	}
	return  res[n]
}
func coins_change_permutat(arr []int , n int) int{
	res := make ([]int , n+1)

	res[0] = 1 
	for i := 1 ;  i < n+1 ; i++{
		temp := 0
		for j := 0 ; j < len(arr) ; j++{
			if i - arr[j] >= 0 {
				temp += res[i - arr[j]]
			}
		}
		res[i] += temp
	}
	return res[n]
}
func main(){
 arr := []int {2,3,5,6}
 n := 8
 fmt.Println(coins_change_comb(arr , n))
 fmt.Println(coins_change_permutat(arr , n))
}