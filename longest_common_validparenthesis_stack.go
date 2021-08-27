type Stack [] int
var top int = -1 

func (s *Stack) isEmpty() bool{
	return len(*s) == 0
}
func (s *Stack) push(value int)  {
	*s = append(*s , value)
	top = top + 1
}
func (s *Stack) pop() (int , bool)  {
	if s.isEmpty() {
		return  0 , false
	}else{
		v :=  (*s)[top]
		*s = (*s)[:top]
		top--
		return v,true		
	}
}
func longestValidParentheses(s string) int {
    
    if len(s) == 0 {
        return 0 
    }
    var stack Stack
    var result int 
    stack.push(-1)
    for i := 0 ; i  < len(s) ; i++{
        
        if s[i] == '('{
            stack.push(i)
        }else{
            if !stack.isEmpty(){
                _,_ = stack.pop()
            }
            if !stack.isEmpty(){
                result = max(result , i - stack[top])
            }else{
                stack.push(i)
            }
        }
    }
    return result
}
func max(a int , b int) int {
    if a > b{
        return a 
    }else{
        return b
    }
}