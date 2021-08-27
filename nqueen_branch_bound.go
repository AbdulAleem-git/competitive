package main

import (
	"fmt"
	"strconv"
)
func solution(board [][]bool , colum , ndiag,revdiag []bool, asf string, row int){
	if row == len(board) {
		fmt.Println(asf+".\n")
		return
	}

	for col := 0 ; col < len(board[0]) ; col++{
		if colum[col] == false && ndiag[row + col] == false && revdiag[row - col + len(board[0]) - 1] == false {
			board[row][col] = true
			colum[col] = true
			ndiag[row+col] = true
			revdiag[row - col + len(board[0])-1] = true
			solution(board , colum , ndiag ,revdiag , asf + strconv.Itoa(row) +"-" + strconv.Itoa(col) + ",", row+1)
			board[row][col] = false
			colum[col] = false
			ndiag[row+col] = false
			revdiag[row - col + len(board[0])-1] = false		
		}
	}
}
func main(){
	board := make([][]bool , 4)
	for i := 0 ; i < len(board) ; i++{
		board[i] = make([]bool , 4)
	}
	colum := make([]bool , len(board))
	ndiag := make([]bool , 2 * len(board) - 1)
	revdiag := make([]bool , 2 * len(board) - 1)
	solution(board , colum , ndiag , revdiag , "" , 0)
}