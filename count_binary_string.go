package main

import "fmt"
func count(n int)int{
	dp_count0 := 1
	dp_count1 := 1
	for i:=2 ; i<=n ; i++{
		new_dp_count0 := dp_count1
		new_dp_count1 := dp_count0 + dp_count1

		dp_count0 = new_dp_count0
		dp_count1 = new_dp_count1
	}
	return dp_count0 + dp_count1
}
func main(){
	fmt.Println(count(7))
	fmt.Println(count(6))
}