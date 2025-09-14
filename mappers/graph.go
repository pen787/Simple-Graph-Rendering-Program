package mappers

import (
	"penrenderingmethod/components"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

func NewGraphMapper(world *ecs.World) *ecs.Map3[components.EmbledScript, rl.Color, string] {
	return ecs.NewMap3[components.EmbledScript, rl.Color, string](world)
}
