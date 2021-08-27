package main

import "fmt"
func printss(s string , ans string){
	if len(s) == 0{
		fmt.Println(ans)
		return
	}
	c := s[0]
	substr := s[1:]
	printss(substr ,""+ans)
	printss(substr ,ans+string(c))
}
func main(){
	printss("abc", "")
}