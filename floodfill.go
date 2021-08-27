package main

import "fmt"

func pathsinmaze(maze [][]int , row , col int , visited[][]bool,ans string){
	if row == len(maze) || col == len(maze[0]) || row < 0 || col < 0 || maze[row][col] == 1|| visited[row][col]==true{
		return
	}
	if row == len(maze)-1 && col == len(maze[0])-1{
		fmt.Println(ans)
		return
	}

	visited[row][col] = true
	pathsinmaze(maze , row-1 , col,  visited, ans + "t")
	pathsinmaze(maze , row , col-1,  visited, ans + "l")
	pathsinmaze(maze , row+1 , col,  visited, ans + "d")
	pathsinmaze(maze , row , col+1,  visited, ans + "r")
	visited[row][col] = false

}

func main(){
	maze := [][]int{
		{0,1,0,0,0,0,0},
		{0,1,0,1,1,1,0},
		{0,0,0,0,0,0,0},
		{1,0,1,1,0,1,1},
		{1,0,1,1,0,1,1},
		{1,0,0,0,0,0,0},
	}
	visited := make([][]bool , len(maze))
	for i:=0 ; i < len(visited) ; i++{
		visited[i] = make([]bool,len(maze[0]))
	}
	pathsinmaze(maze , 0,0,visited ,"")
}