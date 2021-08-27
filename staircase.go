package main

import "fmt"
func staircasecount(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return staircasecount(n-1) + staircasecount(n-2) + staircasecount(n-3)
}
func main(){
	fmt.Println(staircasecount(8))
}