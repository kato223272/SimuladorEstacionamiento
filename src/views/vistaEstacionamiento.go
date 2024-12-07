package views

import (
	"image/color"
	"estacionamiento/src/models"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type EstacionamientoView struct {
	Estacionamiento *models.Estacionamiento
	Context         *scene.Context
}

func NewEstacionamientoView(Estacionamiento *models.Estacionamiento, ctx *scene.Context) *EstacionamientoView {
	pv := &EstacionamientoView{
		Estacionamiento: Estacionamiento,
		Context:         ctx,
	}
	pv.setupScene()
	return pv
}

func (pv *EstacionamientoView) setupScene() {
	ctx := pv.Context

	
	EstacionamientoArea := floatgeom.NewRect2(0, 0, 950, 600)
	backgroundImage, err := render.LoadSprite("assets/Estacionamiento.png")
	if err != nil {
		panic("Error al cargar la imagen de fondo del estacionamiento: " + err.Error())
	}
	entities.New(ctx,
		entities.WithRect(EstacionamientoArea),
		entities.WithRenderable(backgroundImage),
		entities.WithDrawLayers([]int{0}),
	)
	
	
	// entities.New(ctx,
	// 	entities.WithRect(floatgeom.NewRect2(-5, -5, 955, 5)), 
	// 	entities.WithColor(color.RGBA{34, 139, 34, 255}),     
	// 	entities.WithDrawLayers([]int{1}),
	// )
	// entities.New(ctx,
	// 	entities.WithRect(floatgeom.NewRect2(-5, 0, 5, 605)), 
	// 	entities.WithColor(color.RGBA{34, 139, 34, 255}),     
	// 	entities.WithDrawLayers([]int{1}),
	// )
	// entities.New(ctx,
	// 	entities.WithRect(floatgeom.NewRect2(950, 0, 955, 605)), 
	// 	entities.WithColor(color.RGBA{34, 139, 34, 255}),        
	// 	entities.WithDrawLayers([]int{1}),
	// )
	// entities.New(ctx,
	// 	entities.WithRect(floatgeom.NewRect2(-5, 600, 955, 605)), 
	// 	entities.WithColor(color.RGBA{34, 139, 34, 255}),         
	// 	entities.WithDrawLayers([]int{1}),
	// )

	// Fondo de carretera
	roadBackground, err := render.LoadSprite("assets/carretera1.jpg")
	if err != nil {
		panic("Error carretera: " + err.Error())
	}
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(400, 200, 260, 70)),
		entities.WithRenderable(roadBackground),
		entities.WithDrawLayers([]int{0}),
	)

	for _, spot := range pv.Estacionamiento.Spots {
		minX := spot.Area.Min.X()
		minY := spot.Area.Min.Y()
		maxX := spot.Area.Max.X()
		maxY := spot.Area.Max.Y()

		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(minX, minY, maxX, maxY)),
			entities.WithColor(color.RGBA{50, 50, 50, 255}),
			entities.WithDrawLayers([]int{1}),
		)

		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(minX-2, minY-2, maxX+2, minY)), 
			entities.WithColor(color.RGBA{34, 139, 34, 255}),                
			entities.WithDrawLayers([]int{1}),
		)
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(minX-2, maxY, maxX+2, maxY+2)), 
			entities.WithColor(color.RGBA{34, 139, 34, 255}),                   
			entities.WithDrawLayers([]int{1}),
		)

	
		for y := int(minY) + 5; y < int(maxY)-5; y += 15 {
			entities.New(ctx,
				entities.WithRect(floatgeom.NewRect2(minX+5, float64(y), maxX-5, float64(y+2))),
				entities.WithColor(color.RGBA{255, 215, 0, 255}), 
				entities.WithDrawLayers([]int{2}),
			)
		}
	}
		
		// for x := 0; x <= 950; x += 50 { 
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(float64(x), 0, float64(x+30), 15)),
		// 		entities.WithColor(color.RGBA{50, 205, 50, 255}), 
		// 		entities.WithDrawLayers([]int{2}),
		// 	)
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(float64(x), 585, float64(x+30), 600)),
		// 		entities.WithColor(color.RGBA{50, 205, 50, 255}), 
		// 		entities.WithDrawLayers([]int{2}),
		// 	)
		// }
	
		// for y := 0; y <= 600; y += 50 { 
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(0, float64(y), 15, float64(y+30))),
		// 		entities.WithColor(color.RGBA{50, 205, 50, 255}), 
		// 		entities.WithDrawLayers([]int{2}),
		// 	)
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(935, float64(y), 950, float64(y+30))),
		// 		entities.WithColor(color.RGBA{50, 205, 50, 255}), 
		// 		entities.WithDrawLayers([]int{2}),
		// 	)
		// }

		// cornerBushes := []floatgeom.Rect2{
		// 	floatgeom.NewRect2(10, 10, 40, 40),     
		// 	floatgeom.NewRect2(910, 10, 940, 40), 
		// 	floatgeom.NewRect2(10, 560, 40, 590),  
		// 	floatgeom.NewRect2(910, 560, 940, 590), 
		// }
		// for _, bush := range cornerBushes {
		// 	entities.New(ctx,
		// 		entities.WithRect(bush),
		// 		entities.WithColor(color.RGBA{34, 139, 34, 255}), 
		// 		entities.WithDrawLayers([]int{3}),
		// 	)
		// }
		// treePositions := []floatgeom.Point2{
		// 	{50, 50}, {850, 50}, {50, 500}, {850, 500},
		// }
		// for _, pos := range treePositions {
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(pos.X()+10, pos.Y()+20, pos.X()+20, pos.Y()+50)),
		// 		entities.WithColor(color.RGBA{139, 69, 19, 255}), 
		// 		entities.WithDrawLayers([]int{2}),
		// 	)
		// 	entities.New(ctx,
		// 		entities.WithRect(floatgeom.NewRect2(pos.X(), pos.Y(), pos.X()+30, pos.Y()+30)),
		// 		entities.WithColor(color.RGBA{34, 139, 34, 255}), 
		// 		entities.WithDrawLayers([]int{3}),
		// 	)
		// }
	}
