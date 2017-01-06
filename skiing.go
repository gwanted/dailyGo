package main

import "fmt"

type node struct {
	line  int
	visit bool
}

// try to solve skiing problem of DFS
func main() {
	mountain = [5][5]int{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}
	max := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			path[i][j] = dfs(i, j) + 1
			if path[i][j] >= max {
				max = path[i][j]
			}
		}
	}
	fmt.Printf("the longest path size is %d %v\n", max, path)

}

var fuct [5][5]node
var mountain [5][5]int
var path [5][5]int

func dfs(i, j int) int {
	if fuct[i][j].visit {
		return fuct[i][j].line
	}
	fuct[i][j].visit = true
	lens := 0
	tmp := 0
	if i+1 <= 4 && mountain[i+1][j] < mountain[i][j] {
		tmp = dfs(i+1, j) + 1
		fmt.Printf("i+1 %d,j %d,tmp %d\n", i+1, j, tmp)
		if tmp > lens {
			lens = tmp
		}
	}

	if i-1 >= 0 && mountain[i-1][j] < mountain[i][j] {
		tmp = dfs(i-1, j) + 1
		fmt.Printf("i-1 %d,j %d,tmp %d\n", i-1, j, tmp)
		if tmp > lens {
			lens = tmp
		}
	}

	if j+1 <= 4 && mountain[i][j+1] < mountain[i][j] {
		fmt.Printf("i %d,j+1 %d,tmp %d\n", i, j+1, tmp)
		tmp = dfs(i, j+1) + 1
		if tmp > lens {
			lens = tmp
		}
	}

	if j-1 >= 0 && mountain[i][j-1] < mountain[i][j] {
		tmp = dfs(i, j-1) + 1
		fmt.Printf("i %d,j-1 %d,tmp %d\n", i, j-1, tmp)
		if tmp > lens {
			lens = tmp
		}
	}

	fuct[i][j].line = lens
	return lens
}
