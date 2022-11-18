package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets/dirt.png
var dirtBlockPath embed.FS

//go:embed assets/grass.png
var grassBlockPath embed.FS

//go:embed assets/bomb.png
var bombBlockPath embed.FS

//go:embed assets/trap.png
var trapBlockPath embed.FS

//go:embed assets/head.png
var headPath embed.FS

var dirtBlock rl.Texture2D
var grassBlock rl.Texture2D
var bombBlock rl.Texture2D
var trapBlock rl.Texture2D
var headTex rl.Texture2D

func loadAssets() {
	dirtData, _ := dirtBlockPath.ReadFile("assets/dirt.png")
	dirtImage := rl.LoadImageFromMemory(".png", dirtData, int32(len(dirtData)))
	dirtBlock = rl.LoadTextureFromImage(dirtImage)

	grassData, _ := grassBlockPath.ReadFile("assets/grass.png")
	grassImage := rl.LoadImageFromMemory(".png", grassData, int32(len(grassData)))
	grassBlock = rl.LoadTextureFromImage(grassImage)

	bombData, _ := bombBlockPath.ReadFile("assets/bomb.png")
	bombImage := rl.LoadImageFromMemory(".png", bombData, int32(len(bombData)))
	bombBlock = rl.LoadTextureFromImage(bombImage)

	trapData, _ := trapBlockPath.ReadFile("assets/trap.png")
	trapImage := rl.LoadImageFromMemory(".png", trapData, int32(len(trapData)))
	trapBlock = rl.LoadTextureFromImage(trapImage)

	headData, _ := headPath.ReadFile("assets/head.png")
	headImage := rl.LoadImageFromMemory(".png", headData, int32(len(headData)))
	headTex = rl.LoadTextureFromImage(headImage)
}
