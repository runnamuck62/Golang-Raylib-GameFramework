package start

import (
	"GameFrameworkTM/engine"

	ecs "github.com/BrownNPC/simple-ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)
// start scene
type Scene struct {
	pool     *ecs.Pool
	gameOver bool
}

func (scene *Scene) Load(ctx engine.Context) {
	rl.SetTargetFPS(0)
	p := ecs.New(1000)
	scene.pool = p
	spawningSystem(p)
}

func (scene *Scene) Update(ctx engine.Context) (quit bool) {
	p := scene.pool
	if scene.gameOver {
		scene.gameOver = false
		spawningSystem(p)
	}
	renderingSystem(p)

	return false
}

func (scene *Scene) Unload() (nextSceneID string) {
	rl.SetTargetFPS(60)
	return "ohio"
}



