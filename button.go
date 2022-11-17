package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	text               string
	posX               int32
	posY               int32
	width              int32
	height             int32
	textColor          rl.Color
	textColorHighlight rl.Color
	Selected           bool
}

func NewButton(text string, x, y, width, height int32, textColor, textColorHighlight rl.Color) *Button {
	return &Button{
		text:               text,
		posX:               x,
		posY:               y,
		width:              width,
		height:             height,
		textColor:          textColor,
		textColorHighlight: textColorHighlight,
	}
}

func (b *Button) Pressed() bool {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.GetMouseX() > b.posX && rl.GetMouseX() < b.posX+int32(b.width) {
			if rl.GetMouseY() > b.posY && rl.GetMouseY() < b.posY+int32(b.height) {
				return true
			}
		}
	}
	return false
}

func (b *Button) Draw() {
	var fontSize int32 = 20
	textPosX := b.posX + b.width/2 - rl.MeasureText(b.text, fontSize)/2
	textPosY := b.posY + b.height/2 - fontSize/2
	if b.Selected {
		rl.DrawRectangle(b.posX, b.posY, b.width, b.height, rl.NewColor(60, 60, 60, 120))
		rl.DrawText(b.text, textPosX, textPosY, fontSize, b.textColorHighlight)
	} else {
		rl.DrawRectangle(b.posX, b.posY, b.width, b.height, rl.NewColor(60, 60, 60, 120))
		rl.DrawText(b.text, textPosX, textPosY, fontSize, b.textColor)
	}
}
