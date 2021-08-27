package main

import (
	"fmt"
	"strconv"
)

func solution(str ,asf string , count,pos int ){
	if pos == len(str){
		if count  == 0 {
			fmt.Println(asf)
		}else{
			fmt.Println(asf+strconv.Itoa(count))
		}
		return
	} 
	if count > 0 {
		solution(str , asf + strconv.Itoa(count)+string(str[pos]) , 0 , pos+1 )
	}else{
		solution(str, asf + string(str[pos]) , 0 , pos+1 )
	}
	solution(str ,asf , count+1 , pos+1 )
}
func main(){
	solution("abcdefghijklm","",0,0 )
}
