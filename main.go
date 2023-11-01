package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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
}
type GotyaWB struct {
	Hatti1 string `json:"hatti1"`
	Ghoda1 string `json:"ghoda1"`
	Unta1 string `json:"unta1"`
	Vajeer string `json:"vajeer"`
	Raja string `json:"raja"`
	Ghoda2 string `json:"ghoda2"`
	Unta2 string `json:"unta2"`
	Hatti2 string `json:"hatti2"`
	Pyada string `json:"pyada"`
}

func getSructFields(obj interface{}) []string{
	objStruct := reflect.TypeOf(obj)
	if objStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	visible := reflect.VisibleFields(objStruct)
	var res []string
	for _, field :=range visible {
		res = append(res, fmt.Sprintf(field.Name))
	}
	return res
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
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
	gotyaWhiteBlack := getSructFields(GotyaWB{}) 
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
				val:=getAttr(&g.gotya.White,gotyaWhiteBlack[colIndex]) 
				finalVal:= val.Interface().(string)
				g.movePlayer(finalVal, col, gotiColor)

			case 1:	{
				g.movePlayer(g.gotya.White.Pyada, col, gotiColor)
				}
			case 6:	{
				g.movePlayer(g.gotya.White.Pyada, col, gotiColor)
				}

			case 7:
				val:=getAttr(&g.gotya.Black,gotyaWhiteBlack[colIndex]) 
				finalVal:= val.Interface().(string)
				g.movePlayer(finalVal, col, gotiColor)


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
