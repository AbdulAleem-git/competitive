package main

import (
	"fmt"
)

func fib( n int , arr []int) int{
	if n == 1 || n == 0 {
		return n
	}
	if (arr[n] != 0 ){
		return arr[n]
	}
	num := fib(n-1 , arr) + fib(n-2 , arr)
	arr[n] = num
	return num
 }
 func main(){
	for i:=1 ; i <= 20 ; i++{
		arr := make([]int , i+1 )
		fmt.Println(fib(i, arr))
	}
 }