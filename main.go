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
type GotyaEnum int8
const (
	Hatti1 GotyaEnum = iota
	Ghoda1
	Unta1
	Vajeer
	Raja
	Hatti2
	Ghoda2
	Unta2
	Pyada

)

func (g GotyaEnum) String() string {
	switch g{
	case Hatti1: return "hatti1"
	case Ghoda1: return "ghoda1"
	case Unta1: return "unta1"
	case Vajeer: return "vajeer"
	case Raja: return "raja"
	case Hatti2: return "hatti2"
	case Ghoda2: return "ghoda2"
	case Unta2: return "unta2"
	}
	return "pyada"	
}

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
//
func (g Game) movePlayer( goti string) {
		
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
				case 0: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Hatti1, gotiColor, col.color))
					col.player.name = g.gotya.White.Hatti1
					g.movePlayer(g.gotya.White.Hatti1)
					}
				case 1: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Ghoda1, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Ghoda1
					}
				case 2: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Unta1, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Unta1
					}
				case 3: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Vajeer, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Vajeer
					}
				case 4: {g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Raja, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Raja
					}
				case 5: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Unta2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Unta2
					}
				case 6: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Ghoda2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Ghoda2
					}
				case 7: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Hatti2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Hatti2
					}
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
				case 0: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Hatti1, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Hatti1
					}
				case 1: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Ghoda1, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Ghoda1
					}
				case 2: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Unta1, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Unta1
					}
				case 3: {g.level.AddEntity(tl.NewText(x,y, g.gotya.White.Raja, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.White.Raja
					}
				case 4: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Vajeer, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Vajeer
					}
				case 5: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Unta2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Unta2
					}
				case 6: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Ghoda2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Ghoda2
					}
				case 7: {
					g.level.AddEntity(tl.NewText(x,y, g.gotya.Black.Hatti2, gotiColor, col.color))
					g.drawnBoard[rowIndex][colIndex].player.name = g.gotya.Black.Hatti2
					}
				}

			} 
		}
	}
	setInitialPosition(g.drawnBoard)
	g.game.Screen().SetLevel(g.level)
	g.game.Start()
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
