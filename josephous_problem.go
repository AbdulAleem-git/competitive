package main

import "fmt"

func solution(n , k int) int{
	if n == 1 {
		return 0
	}
	x := solution(n-1 , k)
	y := (x + k ) % n
	return y
}
func main(){
	fmt.Println(solution(5 , 3 ))
}