package main

import (
	"fmt"
)

var flag bool = false

func main() {
	maze := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 0, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1}}
	checkpath(maze, 1, 0)
	if !flag {
		fmt.Println("NO")
	}
}
func checkpath(maze [][]int, y int, x int) {
	maze[y][x] = 9
	if y == 7 && x == 8 {
		print(maze, y, x)
		flag = true
	}
	if x <= 7 && maze[y][x+1] == 0 { //右
		checkpath(maze, y, x+1)
	}
	if y <= 7 && maze[y+1][x] == 0 { //下
		checkpath(maze, y+1, x)
	}
	if x > 0 && maze[y][x-1] == 0 { //左
		checkpath(maze, y, x-1)
	}
	if y > 0 && maze[y-1][x] == 0 { //上
		checkpath(maze, y-1, x)
	}
	maze[y][x] = 0
}

func print(maze [][]int, y int, x int) {
	for y = 0; y < len(maze); y++ {
		for x = 0; x < len(maze[0]); x++ {
			if maze[y][x] == 9 {
				fmt.Print("*")
			} else {
				fmt.Print("1")
			}
		}
		fmt.Println()
	}
	fmt.Println("Ok")
}
