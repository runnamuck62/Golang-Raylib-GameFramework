package start

import (
	. "GameFrameworkTM/components" // dot import
	ecs "github.com/BrownNPC/simple-ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func spawningSystem(p *ecs.Pool) {
	ball := ecs.NewEntity(p)
	ecs.Add5(p, ball,
		Position{X: 500, Y: 500},
		Velocity{},
		Ball{},
		Circle(12),
		Color(rl.Red),
	)
	// plate := ecs.NewEntity(p)
	// ecs.Add5(p, plate,
	// 	Position{X: 500, Y: 500},
	// 	Velocity{},
	// 	Plate{},
	// 	Size2D{Width: 20, Height: 20},
	// 	Color(rl.Blue),
	// )
}
func renderingSystem(p *ecs.Pool) {
	POSITION, CIRCLE, SIZE2D, COLOR :=
		ecs.GetStorage4[Position, Circle, Size2D, Color](p)
	for _, e := range POSITION.And(SIZE2D, COLOR) {
		pos, size2d :=
			POSITION.Get(e), SIZE2D.Get(e)
		rl.DrawRectangle(int32(pos.X), int32(pos.Y), size2d.Width, size2d.Height, COLOR.Get(e))
	}
	for _, e := range POSITION.And(CIRCLE, COLOR) {
		pos, circle :=
			POSITION.Get(e), CIRCLE.Get(e)
		rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(circle), COLOR.Get(e))

	}
}
