package main

import (
	"fmt"
	"math"
)

func upstairscountmem(n int , arr[] int)int{
	if n == 0 {
		return 1
	}else if n < 0{
		return 0
	}
	if arr[n] != 0{
		return arr[n]
	}
	r := upstairscountmem(n-1, arr) + upstairscountmem(n-2, arr) + upstairscountmem(n-3 , arr)
	arr[n] = r
	return r
}


func upstairscounttab(n int) int{
	arr := make([]int, n+1)
		for i := 0 ; i < n+1 ; i++{
			if i == 0{
				arr[0] = 1
			}else if i == 1{
				arr[i] = arr[i-1]
			}else if i == 2{
				arr[i] = arr[i-1] + arr[i-2]
			}else{
				arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
			}
		
		}
		return arr[n]
}
func upstairscounttabarr(n int , arr[] int)int{
	res := make([]int , n+1)
	res[n] = 1
	for i :=  n-1 ; i >= 0 ; i--{
		for j := 1 ; j <= arr[i] && i+j < len(res) ; j++{
			res[i] += res[i+j]
		}
	}
	return res[0]
} 
func upstairscountminclimb(n int , arr [] int) int {
	res := make([]int , n+1 )
	for i:=0 ; i < len(res) ; i++{
		res[i] = -1 
	}
	res[n] = 0
	for i := n-1 ; i >= 0 ; i--{
		if arr[i] > 0 {
			temp := int(math.Pow(2 , 32)) - 1
			for j := 1 ; j <= arr[i] && i+j < n+1 ; j++{
				if res[i+j] != -1{
					temp = min(temp , res[i+j])
				}
			}
			if temp != int(math.Pow(2 , 32)) - 1{
				res[i] = temp + 1
			}
		}
		fmt.Printf("i=%d value=%d\n",i,res[i])
	}
	return res[0]
}
func min(a int , b int)int{
	if a < b{
		return a
	}else {
		return b
	}
}
func main(){
	for i:=1 ; i <= 1 ; i++{
		arr := make([]int , i+1 )
		fmt.Println(upstairscountmem(i, arr))
	}
	fmt.Println(upstairscounttab(4))
	arr2 := []int{2,3,1,1,4}
	fmt.Println(upstairscounttabarr(5,arr2))
	fmt.Println(upstairscountminclimb(5,arr2))
 }