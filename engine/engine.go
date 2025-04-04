package engine

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	VirtualWidth, VirtualHeight int
	WindowTitle                 string
}

// map from string id to a Scene
type Scenes map[string]scene

func Run(scenes Scenes, cfg Config) error {
	ActiveSceneId := "start" // look for a scene named start as entry
	ActiveScene, ok := scenes[ActiveSceneId]
	ctx := Context{}
	if !ok {
		return errors.New(`Cannot start. There must be a scene with id "start" that is the entry-point`)
	} else if ActiveScene == nil {
		return errors.New("start scene cannot be nil")
	}
	// -----------------------------BEGIN------------------------------------
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(0, 0, cfg.WindowTitle)
	rl.InitAudioDevice()
	defer rl.CloseWindow() // de-initialization
	defer rl.CloseAudioDevice()
	// -----------------------CENTER WINDOW----------------------------
	WindowWidth, WindowHeight := (rl.GetScreenWidth()*90)/100, (rl.GetScreenHeight()*90)/100
	rl.SetWindowSize(WindowWidth, WindowHeight) //90% of screen
	centerWindow()
	// ----LOAD START SCENE----
	ActiveScene.Load(ctx)
	// ----MAIN LOOP----
	for !rl.WindowShouldClose() {
		// ----FULL SCREEN ON F11----
		if rl.IsKeyPressed(rl.KeyF11) {
			rl.ToggleBorderlessWindowed()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// -------UPDATE SCENE---------
		quit := ActiveScene.Update(ctx)
		rl.EndDrawing()

		if quit {
			// -------UNLOAD SCENE-------
			nextSceneId := ActiveScene.Unload()
			nextScene, ok := scenes[nextSceneId]
			// ------SWITCH SCENE------
			if ok && nextScene != nil {
				ActiveSceneId = nextSceneId
				ActiveScene = nextScene
				ActiveScene.Load(ctx)
				continue
			}
			//-----ERROR HANDLING------
			if !ok {
				return fmt.Errorf(`There is no scene with id "%s", tried switching from scene "%s"`, nextSceneId, ActiveSceneId)
			} else if nextScene == nil {
				return fmt.Errorf(`scene with id "%s" is nil, tried switching from scene "%s"`, nextSceneId, ActiveSceneId)
			}
		}
	}
	return nil
}
func centerWindow() {
	WindowWidth, WindowHeight := rl.GetScreenWidth(), rl.GetScreenHeight()
	monitorWidth := rl.GetMonitorWidth(rl.GetCurrentMonitor())
	monitorHeight := rl.GetMonitorHeight(rl.GetCurrentMonitor())
	xPos := (monitorWidth - WindowWidth) / 2
	yPos := (monitorHeight - WindowHeight) / 2
	rl.SetWindowPosition(xPos, yPos)
}
