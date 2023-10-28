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
	X int
	Y int
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
	gotiw WGotyaEnum
	gotib BGotyaEnum
}
type BGotyaEnum int16
const (
	BHatti1 BGotyaEnum = 1 << iota
	BGhoda1
	BUnta1
	BVajeer
	BRaja
	BHatti2
	BGhoda2
	BUnta2
	BPyada

)

func (g BGotyaEnum) String() string {
	switch g{
	case BHatti1: return "bhatti1"
	case BGhoda1: return "bghoda1"
	case BUnta1: return "bunta1"
	case BVajeer: return "bvajeer"
	case BRaja: return "braja"
	case BHatti2: return "bhatti2"
	case BGhoda2: return "bghoda2"
	case BUnta2: return "bunta2"
	}
	return "wpyada"	
}

type WGotyaEnum int16
const (
	WHatti1 WGotyaEnum = 1 << iota
	WGhoda1
	WUnta1
	WVajeer
	WRaja
	WHatti2
	WGhoda2
	WUnta2
	WPyada

)

func (g WGotyaEnum) String() string {
	switch g{
	case WHatti1: return "whatti1"
	case WGhoda1: return "wghoda1"
	case WUnta1: return "wunta1"
	case WVajeer: return "wvajeer"
	case WRaja: return "wraja"
	case WHatti2: return "whatti2"
	case WGhoda2: return "wghoda2"
	case WUnta2: return "wunta2"
	}
	return "wpyada"	
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
func (g Game) movePlayer(goti string, currBlock block, goticolor tl.Attr) {
	g.drawnBoard[currBlock.row][currBlock.col].player.name = goti
	x:=g.drawnBoard[currBlock.row][currBlock.col].X 
	y:=g.drawnBoard[currBlock.row][currBlock.col].Y 
	g.level.AddEntity(tl.NewText(x,y, goti, goticolor, currBlock.color))


}
func (g Game) paintBoard() {
	for rowIndex, row := range g.drawnBoard {
		y:= rowIndex * 6
		for colIndex, col:= range row {
			x:= colIndex * 12
			g.level.AddEntity(tl.NewRectangle(x, y, 12, 6, col.color))
			g.drawnBoard[rowIndex][colIndex].X = x
			g.drawnBoard[rowIndex][colIndex].Y = y
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
					g.movePlayer(g.gotya.White.Hatti1, col, gotiColor)
					}
				case 1: {
					g.movePlayer(g.gotya.White.Ghoda1, col, gotiColor)
					}
				case 2: {
					g.movePlayer(g.gotya.White.Unta1, col, gotiColor)
					}
				case 3: {
					g.movePlayer(g.gotya.White.Vajeer, col, gotiColor)
					}
				case 4: {
					g.movePlayer(g.gotya.White.Raja, col, gotiColor)
					}
				case 5: {
					g.movePlayer(g.gotya.White.Unta2, col, gotiColor)
					}
				case 6: {
					g.movePlayer(g.gotya.White.Ghoda2, col, gotiColor)
					}
				case 7: {
					g.movePlayer(g.gotya.White.Hatti2, col, gotiColor)
					}
				}
			case 1:	{
				g.movePlayer(g.gotya.White.Pyada, col, gotiColor)
			}
			case 6:	{
				g.movePlayer(g.gotya.White.Pyada, col, gotiColor)
				}
				
			case 7:
				switch colIndex {
				case 0: {
					g.movePlayer(g.gotya.Black.Hatti1, col, gotiColor)
					}
				case 1: {
					g.movePlayer(g.gotya.Black.Ghoda1, col, gotiColor)
					}
				case 2: {
					g.movePlayer(g.gotya.Black.Unta1, col, gotiColor)
					}
				case 3: {
					g.movePlayer(g.gotya.Black.Raja, col, gotiColor)
					}
				case 4: {
					g.movePlayer(g.gotya.Black.Vajeer, col, gotiColor)
					}
				case 5: {
					g.movePlayer(g.gotya.Black.Unta2, col, gotiColor)
					}
				case 6: {
					g.movePlayer(g.gotya.Black.Ghoda2, col, gotiColor)
					}
				case 7: {
					g.movePlayer(g.gotya.Black.Hatti2, col, gotiColor)
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
	var gotyaW WGotyaEnum
	var gotyaB BGotyaEnum
	for i:=WHatti1; i <= WPyada; i= i<<1 {
		fmt.Println("does this even works", i)
	}
	g := Game{
		drawnBoard: &drawnBorad,
		game: tl.NewGame(),
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlue,
			Ch: '/',
		}),
		gotya: gotya,
		gotiw: gotyaW,
		gotib: gotyaB,
	}	
	g.paintBoard()
}
