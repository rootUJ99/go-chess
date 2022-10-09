package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	fmt.Println("hello")
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Ch: 'O',
	})
	hrPtr := 0
	vrPtr := 0
	for i := 0; i <= 63; i++ {
		block := i + 1
		x := vrPtr * 12
		y := hrPtr * 6
		vrPtr += 1
		color := tl.ColorBlack

		if int(hrPtr%2) == 0 {
			color = tl.ColorWhite
			if int(vrPtr%2) == 0 {
				color = tl.ColorBlack
			}
		} else {
			color = tl.ColorBlack
			if int(vrPtr%2) == 0 {
				color = tl.ColorWhite
			}
		}
		if block%8 == 0 {
			hrPtr += 1
			vrPtr = 0
		}

		level.AddEntity(tl.NewRectangle(x, y, 12, 6, color))
		game.Screen().SetLevel(level)
	}
	game.Start()
}
