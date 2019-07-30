package main

import (
	"fmt"
)

var flag bool

type point struct {
	y, x int
}

//方向、上左下右
var direction = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.y + r.y, p.x + r.x}
}

//outboundのcheck
func (p point) outbound(maze [][]int) (int, bool) {
	if p.y < 0 || p.y >= len(maze) {
		return 0, false
	}
	if p.x < 0 || p.x >= len(maze[p.y]) {
		return 0, false
	}

	return maze[p.y][p.x], true
}

//通路探す。
func walk(maze [][]int, start, end point) [][]int {
	flag = false
	newmaze := make([][]int, len(maze))
	for i := range newmaze {
		newmaze[i] = make([]int, len(maze[i]))
	}

	//チェックのキュー
	Q := []point{start}

	for len(Q) > 0 {
		check := Q[0] //point{1,0}
		Q = Q[1:]     //値なし。

		if check == end {
			fmt.Println("  ", "通路見つかりました。")
			flag = true
			break
		}
		//通路探す処理
		for _, dir := range direction {
			next := check.add(dir) //出した場所のチェック

			val, ok := next.outbound(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.outbound(newmaze)
			if !ok || val != 0 {
				continue
			}
			//最初インクリメントしていない為、戻ってはいけない。
			if next == start {
				continue
			}
			//この時点のstep数
			nowstep, _ := check.outbound(newmaze)
			newmaze[next.y][next.x] = nowstep + 1 //探索した所、インクリメントして、新しいの地図に入れる。
			//nextを次探索所として、キューに入れる。
			Q = append(Q, next)
		}
	}
	return newmaze
}

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

	newmaze := walk(maze, point{1, 0}, point{7, 8})

	for _, y := range newmaze {
		for _, val := range y {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	if !flag {
		fmt.Println("通路がない!!")
	}
}
