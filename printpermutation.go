package main

import "fmt"

func printpermutation(s string , ans string){
	if len(s)==0 && len(ans)==0{
		return 
	}else if len(s)==0{
		fmt.Print(ans+" ")
		return
	}
	for i := 0 ; i < len(s) ; i++{
		printpermutation(s[:i]+s[i+1:] , ans+string(s[i]))
	}
}
func main(){
	printpermutation("abcd","")
}