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
	drawnBoard board
	game *tl.Game
	level *tl.BaseLevel
}

func (g Game) constructBoard() board{
	for rowIndex, row := range g.drawnBoard{
		for colIndex, _ := range row{
			if (colIndex+1+rowIndex) % 2 == 0{
				g.drawnBoard[rowIndex][colIndex].color = tl.ColorBlack
			} else {
				g.drawnBoard[rowIndex][colIndex].color = tl.ColorWhite
			}
			g.drawnBoard[rowIndex][colIndex].row = rowIndex
			g.drawnBoard[rowIndex][colIndex].col = colIndex
		}
	}
	return g.drawnBoard
}

func (g Game) paintBoard(gboard board) {
	for rowIndex, row := range gboard {
		y:= rowIndex * 6
		for colIndex, col:= range row {
		 x:= colIndex * 12
			g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, col.color))
		} 
	}
	g.game.Screen().SetLevel(g.level)

}


func main() {
	fmt.Println("Starting the game!!")
	g := Game{
		game: tl.NewGame(),
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlue,
			Ch: '/',
		}),
	}	
	gboard:=g.constructBoard()
	g.paintBoard(gboard)
	g.game.Start()
}
