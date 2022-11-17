# MineCreeper
A Minesweeper clone in the Minecraft universe written in go and raylib

![](readme/screenshot001.png)
![](readme/screenshot002.png)
![](readme/screenshot003.png)
![](readme/screenshot004.png)

## Requirements
- [raylib-go](https://github.com/gen2brain/raylib-go) see the [requirements](https://github.com/gen2brain/raylib-go#requirements) of `raylib-go`

## How to run
- run `go get github.com/gen2brain/raylib-go/raylib` to get raylib-go
- run `go run .` within the root folder

## Build release
- run `go build -ldflags "-s -w"`
> Note: To prevent the console window from popping up on Windows add `-H windowsgui` to the linker flags -> `go build -ldflags "-s -w -H windowsgui" .`

> Note: If you run into issues distributing the game on windows (`libwinpthread-1.dll not found`) try to build with CGO_LDFLAGS to statically link the files: `$env:CGO_LDFLAGS = "-static-libgcc -static -lpthread"; go build -ldflags "-s -w -H windowsgui" .`

Issue has been discussed here: [raylib-go issues](https://github.com/gen2brain/raylib-go/issues/135#issuecomment-895562752)