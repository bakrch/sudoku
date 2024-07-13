package main

type Grid [9][9]int

type Sudoku struct {
	grid Grid
}

func NewSudoku(grid [9][9]int) Sudoku {
	return Sudoku{
		grid: grid,
	}
}

func (p *Sudoku) Check() {

}
func (p *Sudoku) Solve() {

}

func main() {
	newGrid := [9][9]int{
		{0, 0, 0, 1, 5, 0, 4, 0, 0},
		{6, 0, 9, 2, 8, 4, 0, 0, 0},
		{0, 0, 0, 3, 9, 6, 0, 1, 8},
		{3, 0, 1, 5, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 3, 0, 6, 0, 9},
		{4, 9, 8, 6, 0, 2, 0, 0, 1},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},
		{9, 4, 0, 7, 6, 3, 0, 0, 5},
		{0, 3, 7, 0, 1, 0, 0, 6, 0},
	}
	sudoku := NewSudoku(newGrid)
	sudoku.Check()
}
