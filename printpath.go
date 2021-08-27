package main

import (
	"fmt"
	"strconv"
)
func  printpath(n int , ans string)  {
	if n <= 0 {
		if n < 0 {
			return 
		}else{
			fmt.Println(ans)
			return
		}
	}
	ans =   ans+strconv.Itoa(n)
	printpath(n-1, ans) 
	printpath(n-2, ans) 
	printpath(n-3, ans) 
}
func main(){
	printpath(4, "")
}