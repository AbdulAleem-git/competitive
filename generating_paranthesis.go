package main

import "fmt"
func generateParenthesis(n int) []string {
    
    res := make(map[int][]string)
	res[0] = []string{""}
	res[1] = []string{"()"}
	for i := 2 ; i <= n ; i++{
		for j ,k:= i-1,0 ; k <= i-1 && j>=0 ; j,k = j-1,k+1 {
			str_j := res[j]
			str_k := res[k]
			temp_str_j := make([]string , len(str_j))
			for l := 0 ; l < len(str_j) ; l++{
				temp_str_j[l] = "(" + str_j[l] + ")"
			}
			
			for l := 0 ; l < len(temp_str_j) ; l++{
				for m:=0 ; m< len(str_k); m++{
					res[i] = append(res[i], temp_str_j[l]+str_k[m])
				}
			}
		}
	}
	return res[n]
}
func main()  {
	fmt.Println(generateParenthesis(8))
}