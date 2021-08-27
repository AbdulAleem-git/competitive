package main

import "fmt"
func waystopaint(color int , pencs int)int{
	if pencs >=2 {
		same := color
		diff := color * (color-1)
		total := same + diff
		for i := 3 ; i <= pencs ; i++{
			same = diff
			diff = total 
			total = same + diff
		}
		return total
	}else{
		return 0 
	}
}
func main(){
	fmt.Println(waystopaint(3,5))
}