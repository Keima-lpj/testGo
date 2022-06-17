package main

import (
	"fmt"
	"os"
)

//迷宫的广度优先算法
func main() {
	maze := readMaze("maze/maze")
	step := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for i := range step {
		for j := range step[i] {
			fmt.Printf("%d ", step[i][j])
		}
		fmt.Println()
	}
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//这个方法用于判断当前的点在二位数组中的位置
func (p point) at(grid [][]int) (int, bool) {
	//判断p是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//这个代表着当前探索的点要走的四个方向
var dirs = []point{
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

//开始走迷宫
func walk(maze [][]int, start, end point) [][]int {
	//新建一个slice，用来存储走过的路径
	step := make([][]int, len(maze))
	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}

	//一个队列，存入将要探索的点
	queue := []point{start}
	//当队列中还有没有探索的值的时候
	for len(queue) > 0 {
		//将队列的第一个值弹出来，开始探索它
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		//循环四个方向
		for _, dir := range dirs {
			next := cur.add(dir)
			//判断三个条件
			//1、当前点不能等于起点，如果等于起点，则直接跳过循环
			if next == start {
				continue
			}
			//2、当前点在迷宫中不能是1并且不能越界
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			//3、当前点在step中不能有值，必须是0（因为有值代表已经走过了）
			val, ok = next.at(step)
			if !ok || val != 0 {
				continue
			}
			//将这个点放入step中
			//当前点的步骤数
			val, _ = cur.at(step)
			step[next.i][next.j] = val + 1
			//将这个点放入队列中以后再探索
			queue = append(queue, next)
		}
	}

	return step
}

//从文件中读取迷宫
func readMaze(filename string) [][]int {
	//定义行和列
	var row, col int
	//打开并读取文件
	file, err := os.Open(filename)
	if err != nil {
		panic(nil)
	}
	//读取文件的第一行
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
