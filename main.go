package main

import (
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {
	rl.SetTraceLogLevel(rl.LogError)
	err := engine.Run(scenes.Registered, engine.Config{
		WindowTitle: "My gyatt",
	})
	if err != nil {
		fmt.Println(err)
	}
}
