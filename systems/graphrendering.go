package systems

import (
	"log"
	"penrenderingmethod/components"
	"penrenderingmethod/utility"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

type GraphRendering struct {
	world        *ecs.World
	filter       *ecs.Filter3[components.EmbledScript, rl.Color, string]
	cameraOrigin *rl.Vector2
	zoom         *float32
}

func (r *GraphRendering) Init() {
	query := r.filter.Query()
	for query.Next() {
		es, gotcolor, path := query.Get()
		err := es.DoFile(*path)
		if err != nil {
			log.Fatalln(err)
		}

		new_color, err := es.CallLoad()
		if err != nil {
			log.Fatalln(err)
		}
		gotcolor.R = new_color.R
		gotcolor.G = new_color.G
		gotcolor.B = new_color.B
		gotcolor.A = new_color.A
	}

}

func (r *GraphRendering) Update() {}

func (r *GraphRendering) Render() {
	query := r.filter.Query()
	for query.Next() {
		es, retcolor, _ := query.Get()

		utility.DrawGraph(2, func(x float32) float32 {
			ret, err := es.CallRender(x)
			if err != nil {
				log.Fatalln(err)
			}
			return ret
		}, *retcolor, *r.cameraOrigin, *r.zoom)
	}
}

func (r *GraphRendering) Done() {
	query := r.filter.Query()
	for query.Next() {
		es, _, _ := query.Get()
		es.Close()
	}
}

func NewRendering(world *ecs.World, cameraOrigin *rl.Vector2, Zoom *float32) *GraphRendering {
	return &GraphRendering{
		world:        world,
		filter:       ecs.NewFilter3[components.EmbledScript, rl.Color, string](world),
		cameraOrigin: cameraOrigin,
		zoom:         Zoom,
	}
}
