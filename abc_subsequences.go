package main

import "fmt"
func subsequences(str string)int {
	a := 0 
	ab := 0
	abc := 0
	for i := 0 ; i < len(str) ; i++{
		if str[i] == 'a'{
			a = 2 * a + 1
		}else if str[i] == 'b'{
			ab = 2 * ab  + a 
		}else if str[i] == 'c'{
			abc = 2 * abc + ab
		}
	}
	return abc
}
//A+B+C
func main(){
	fmt.Println(subsequences("abcabc"))
}

// cpalgorithm.com
// practice
// contest
// 2thousand
// dsa learning
// a2ojladder
