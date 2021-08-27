package main

import "fmt"

func fibnacci_iterative( n int )int{
	firstp := 0 
	secondp := 1
	for i := 0 ; i < n - 1; i++{
		current := firstp + secondp
		firstp = secondp
		secondp = current
	}
	return secondp
}
func fibnacci_rec(n int , asf [] int)int{
	if n == 0 || n == 1 {
		return n
	}
	if asf[n] != 0{
		return asf[n]
	}
	fmt.Println("hello ",n)
	fn :=  fibnacci_rec(n-1,asf) + fibnacci_rec(n-2,asf)
	asf[n] = fn
	return fn
} 

func main(){
	n := 10 
	asf := make ([]int , n+1)
	fmt.Println(fibnacci_iterative(10))
	fmt.Println(fibnacci_rec(10 , asf))
}