package main

import "fmt"

type Stack [] string

func (s *Stack) isEmpty() bool{
	return len(*s) == 0
}
func (s *Stack) push(value string)  {
	*s = append(*s , value)
}
func (s *Stack) pop() (string , bool)  {
	if s.isEmpty() {
		return  "" , false
	}else{
		v :=  (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return v,true		
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	test := "hello world"
	for _,i := range test{
		stack.push(string(i))
	}
	for len(stack) > 0 {
		x, y := stack.pop()
		if y == true {
			fmt.Println(x)
		}
	}
}