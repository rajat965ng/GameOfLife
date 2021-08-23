package service

import "fmt"

type gameOfLife struct{}

func NewInstance() *gameOfLife {
	return &gameOfLife{}
}

func (gol *gameOfLife) Display(arr [][]int) {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			switch arr[i][j] {
			case 1:
				fmt.Print("*")
			default:
				fmt.Print(".")

			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (gol *gameOfLife) NextGenerate(arr [][]int, n int, m int) [][]int {

	next := make([][]int, n)
	for k := range next {
		next[k] = make([]int, m)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			alive := 0

			for x := -1; x <= 1; x++ {
				for z := -1; z <= 1; z++ {
					alive += arr[i+x][j+z]
				}
			}
			alive -= arr[i][j]

			switch {
			case arr[i][j] == 1 && alive > 3:
				next[i][j] = 0
			case arr[i][j] == 1 && alive < 2:
				next[i][j] = 0
			case arr[i][j] == 0 && alive == 3:
				next[i][j] = 1
			default:
				next[i][j] = arr[i][j]
			}
		}
	}
	return next
}
