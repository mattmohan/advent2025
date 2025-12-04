package utils

import "strings"

type Grid[T any] struct {
	field []T
	rows  int
	cols  int
}

// NewGrid creates a new Grid with the specified number of rows and columns, initialized with the defaultValue.
func NewGrid[T any](rows, cols int, defaultValue T) Grid[T] {
	field := make([]T, rows*cols)
	for i := range field {
		field[i] = defaultValue
	}
	return Grid[T]{field: field, rows: rows, cols: cols}
}

// Get retrieves the value at the specified (x, y) coordinates in the grid.
func (g Grid[T]) Get(x, y int) T {
	if y < 0 || y >= g.rows || x < 0 || x >= g.cols {
		panic("Index out of bounds")
	}
	return g.field[y*g.cols+x]
}

// Set sets the value at the specified (x, y) coordinates in the grid.
func (g *Grid[T]) Set(x, y int, value T) {
	if y < 0 || y >= g.rows || x < 0 || x >= g.cols {
		panic("Index out of bounds")
	}
	g.field[y*g.cols+x] = value
}

// Rows returns the number of rows in the grid.
func (g Grid[T]) Rows() int {
	return g.rows
}

// Cols returns the number of columns in the grid.
func (g Grid[T]) Cols() int {
	return g.cols
}

// Clones the contents of grid2 into the current grid.
func (g *Grid[T]) Clone(grid2 Grid[T]) {
	copy(g.field, grid2.field)
}

// Walks through each cell in the grid and applies the provided function.
func (g Grid[T]) Walk(f func(x, y int, value T)) {
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			f(x, y, g.Get(x, y))
		}
	}
}

// WalkNeighbors walks through the neighboring cells of (x, y) and applies the provided function.
func (g Grid[T]) WalkNeighbors(x, y int, f func(x, y int, value T)) {
	for row := y - 1; row <= y+1; row++ {
		for col := x - 1; col <= x+1; col++ {
			if row == y && col == x {
				continue
			}
			if row < 0 || row >= g.rows || col < 0 || col >= g.cols {
				continue
			}
			f(col, row, g.Get(col, row))
		}
	}
}

// String returns a string representation of the grid using the provided valueToString function.
func (g Grid[T]) String(valueToString func(value T) rune) string {
	var sb strings.Builder
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			sb.WriteRune(valueToString(g.Get(x, y)))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
