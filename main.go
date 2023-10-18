package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)


type board [8][8]block

type block struct {
	color tl.Attr 
	player player
	row int
	col int
}

type player struct {
	color tl.Attr 
	name string
}

type Game struct {
	drawnBoard *board
	game *tl.Game
	level *tl.BaseLevel
}

func constructBoard() board{
	var drawnBoard board
	for rowIndex, row := range drawnBoard{
		for colIndex, _ := range row{
			if (colIndex+1+rowIndex) % 2 == 0{
				drawnBoard[rowIndex][colIndex].color = tl.ColorBlack
			} else {
				drawnBoard[rowIndex][colIndex].color = tl.ColorWhite
			}
			drawnBoard[rowIndex][colIndex].row = rowIndex
			drawnBoard[rowIndex][colIndex].col = colIndex
		}
	}
	return drawnBoard
}

func (g Game) paintBoard() {
	for rowIndex, row := range g.drawnBoard {
		y:= rowIndex * 6
		for colIndex, col:= range row {
		 x:= colIndex * 12
			g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, col.color))
		} 
	}
	g.game.Screen().SetLevel(g.level)
	g.game.Start()
}


func main() {
	fmt.Println("Starting the game!!")
	drawnBorad:= constructBoard()
	g := Game{
		drawnBoard: &drawnBorad,
		game: tl.NewGame(),
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlue,
			Ch: '/',
		}),
	}	
	g.paintBoard()
}
