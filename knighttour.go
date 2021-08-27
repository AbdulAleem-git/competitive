package main

import "fmt"
func tour(chess [][]int , row , col , move int){
	if row < 0 || col < 0 || row >= len(chess) || col >= len(chess[0]) || chess[row][col] > 0 {
		return
	}else if (move == (len(chess) * len(chess[0]) )){
		chess[row][col] = move
		displaychessboard(chess)
		chess[row][col] = 0
		return
	}

	chess[row][col] = move
	tour(chess , row-2 , col+1 , move+1 )
	tour(chess , row-1 , col+2 , move+1 )
	tour(chess , row+1 , col+2 , move+1 )
	tour(chess , row+2 , col+1 , move+1 )
	tour(chess , row+2 , col-1 , move+1 )
	tour(chess , row+1 , col-2 , move+1 )
	tour(chess , row-1 , col-2 , move+1 )
	tour(chess , row-2 , col-1 , move+1 )
	chess[row][col] = 0
}
func displaychessboard(chess [][]int){
	for i := 0 ; i < len (chess) ; i++{
		fmt.Print("{")
		for j:=0 ; j < len(chess[0]);j++{
			fmt.Printf("%d ", chess[i][j])
		}
		fmt.Printf("}\n")
	}
	fmt.Printf("\n")
}
func main(){
	chess := make([][]int , 5)
	for i ,_ := range chess{
		chess[i] = make([]int , 5)
	}
	tour(chess , 3 ,3,1)
}