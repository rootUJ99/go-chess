package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	gotya Gotya
}

const (
	Hatti1="hatti1"
	Ghoda1="ghoda1"
	Unta1="unta1"
	Vajeer="vajeer"
	Raja="raja"
	Hatti2="hatti2"
	Ghoda2="ghoda2"
	Unta2="unta2"
	Pyada="pyada"

)

type GotyaWB struct {
	Hatti1 string `json:"hatti1"`
	Ghoda1 string `json:"ghoda1"`
	Unta1 string `json:"unta1"`
	Vajeer string `json:"vajeer"`
	Raja string `json:"raja"`
	Hatti2 string `json:"hatti2"`
	Ghoda2 string `json:"ghoda2"`
	Unta2 string `json:"unta2"`
	Pyada string `json:"pyada"`
}

type Gotya struct {
	White GotyaWB `json:"white"` 
	Black GotyaWB `json:"black"` 
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

func setInitialPosition(drawnBoard *board) {
	for _, row := range drawnBoard {
		fmt.Println(row)
	}
}

func (g Game) paintBoard() {
	for rowIndex, row := range g.drawnBoard {
		y:= rowIndex * 6
		for colIndex, col:= range row {
			x:= colIndex * 12
			g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, col.color))
			var gotiColor tl.Attr
			if col.color == tl.ColorBlack {
				gotiColor = tl.ColorWhite
			} else {
				gotiColor = tl.ColorBlack
			}
			switch rowIndex {
			case 0:
				switch colIndex {
				case 0: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Hatti1, gotiColor, col.color))
				case 1: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Ghoda1, gotiColor, col.color))
				case 2: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Unta1, gotiColor, col.color))
				case 3: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Vajeer, gotiColor, col.color))
				case 4: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Raja, gotiColor, col.color))
				case 5: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Unta2, gotiColor, col.color))
				case 6: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Ghoda2, gotiColor, col.color))
				case 7: g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Hatti2, gotiColor, col.color))
				}
			case 1:	{
				g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Pyada, gotiColor, col.color)) 
				g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Pyada
			}
			case 6:	{
				g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Pyada, gotiColor, col.color))
				g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Pyada
				}
				
			case 7:
				switch colIndex {
				case 0: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Hatti1, gotiColor, col.color))
				case 1: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Ghoda1, gotiColor, col.color))
				case 2: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Unta1, gotiColor, col.color))
				case 3: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Raja, gotiColor, col.color))
				case 4: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Vajeer, gotiColor, col.color))
				case 5: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Unta2, gotiColor, col.color))
				case 6: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Ghoda2, gotiColor, col.color))
				case 7: g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Hatti2, gotiColor, col.color))
				}

			} 
		}
	}
	g.game.Screen().SetLevel(g.level)
	g.game.Start()
	setInitialPosition(g.drawnBoard)
}


func main() {
	fmt.Println("Starting the game!!")
	drawnBorad:= constructBoard()
	file, _ := os.ReadFile("./gotya.json")
	var gotya Gotya
	err:= json.Unmarshal(file, &gotya)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gotya)
	g := Game{
		drawnBoard: &drawnBorad,
		game: tl.NewGame(),
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlue,
			Ch: '/',
		}),
		gotya: gotya,
	}	
	g.paintBoard()
}
