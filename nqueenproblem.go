package main

import (
	"fmt"
	"strconv"
)

func printqueens(chess [][]int, psf string, row int) {
	if len(chess) == row{
		fmt.Println(psf+".")
		return
	}
	for col := 0; col < len(chess[0]); col++ {
		if isqueensafe(chess , row , col) {
			chess[row][col] = 1
			printqueens(chess, psf+strconv.Itoa(row)+"-"+strconv.Itoa(col)+"," , row+1)
			chess[row][col] = 0
		}
	}
}
func isqueensafe(chess[][]int , row , col int)bool{
	for i , j := row - 1 ,  col ; i >= 0; i--{
		if chess[i][j] == 1 {
			return false
		}
	}

	for i , j := row -1 , col -1 ; i >= 0 && j >=0 ; i,j = i-1 , j-1 {
		if chess[i][j] == 1 {
			return false
		}
	}
	for i , j := row -1 , col +1 ; i >= 0 && j < len(chess[0]) ; i,j = i-1 , j+1 {
		if chess[i][j] == 1 {
			return false
		}
	}
	return true
}
func main(){
	chess := make([][]int , 5)
	for i := 0 ; i < len(chess); i++{
		chess[i] = make([]int , 5)
	}
	printqueens(chess , "", 0)
}