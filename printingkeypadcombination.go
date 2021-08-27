package main

import (
	"fmt"
	"strconv"
)


func printcombination(str string , keypad []string)[]string{
	var res [] string
	if str == "" {
		res = append(res, "")
		return res			
	}
	number ,_:= strconv.Atoi(string(str[0]))
	temp := printcombination(str[1:] ,keypad)
	for i := 0 ; i < len(keypad[number]) ; i++{
		for j := 0 ; j < len(temp) ; j++{
			res = append(res, string(keypad[number][i]) + temp[j])
		}

	}
	return res
}
func main(){
	keypad := []string{"?!","abc","def","ghi","jkl","mnop","qrst","uv","wxyz","%"}
	fmt.Println(printcombination("678",keypad))
	
}