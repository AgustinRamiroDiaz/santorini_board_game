package main

import (
	"fmt"
)

type worker struct{}

type cell struct {
	blocksBuilt uint
	hasDome     bool
	worker      *worker
}

type board [5][5]cell

func (board board) build(x, y uint) {
	board[x][y].blocksBuilt++
}

func (board board) buildDome(x, y uint) {
	board[x][y].hasDome = true
}

func (board board) placeWorker(x, y uint, worker *worker) {
	board[x][y].worker = worker
}

func (board board) moveWorker(xFrom, yFrom, xTo, yTo uint) {
	board[xTo][yTo].worker = board[xFrom][yFrom].worker
	board[xFrom][yFrom].worker = nil
}

func (board board) canMove(xFrom, yFrom, xTo, yTo uint) bool {
	cellFrom := board[xFrom][yFrom]
	cellTo := board[xTo][yTo]
	return cellTo.worker == nil &&
		cellTo.blocksBuilt-cellFrom.blocksBuilt <= 1 &&
		!cellTo.hasDome
}

func main() {
	fmt.Println("Hello, World!")
}
