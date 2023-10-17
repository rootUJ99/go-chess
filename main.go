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

func (g Game) paintBoard() {
	for rowIndex, row := range g.drawnBoard {
		/* y:= rowIndex * 6 */
		y:= rowIndex * 12
		for colIndex, col:= range row {
			/* x:= colIndex * 12 */
			x:= colIndex * 6
			g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, col.color))
			g.game.Screen().SetLevel(g.level)
		} 
	}
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
	g.constructBoard()
	g.paintBoard()
	// fmt.Println(reflect.TypeOf(tl.ColorBlack))
	// hrPtr := 0
	// vrPtr := 0
	// for i := 0; i <= 63; i++ {
	// 	block := i + 1
	// 	x := vrPtr * 12
	// 	y := hrPtr * 6
	// 	vrPtr += 1
	// 	color := tl.ColorBlack
	//
	// 	if int(hrPtr%2) == 0 {
	// 		color = tl.ColorWhite
	// 		if int(vrPtr%2) == 0 {
	// 			color = tl.ColorBlack
	// 		}
	// 	} else {
	// 		color = tl.ColorBlack
	// 		if int(vrPtr%2) == 0 {
	// 			color = tl.ColorWhite
	// 		}
	// 	}
	// 	if block%8 == 0 {
	// 		hrPtr += 1
	// 		vrPtr = 0
	// 	}
	//
	// 	fmt.Println(x,y)
	// 	g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, color))
	// 	g.game.Screen().SetLevel(g.level)
	// }
	fmt.Println(len(g.level.Entities))
	g.game.Start()
}
