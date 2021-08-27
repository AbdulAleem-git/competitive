package main

import "fmt"

func arrange_building(n int) int {
	dpc0 := 1
	dpc1 := 1
	for i := 2 ; i <=n ; i++{
		dpc0_new := dpc1
		dpc1_new := dpc0 + dpc1
		
		dpc0 = dpc0_new
		dpc1 = dpc1_new
	}
	return (dpc0 + dpc1) * (dpc0 + dpc1)
}
func main(){
	fmt.Println(arrange_building(7))
}