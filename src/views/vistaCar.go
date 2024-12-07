package views

import (
	"estacionamiento/src/models"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)
type CarView struct {
	Car     *models.Car
	Sprite  *render.Switch
	Context *scene.Context
}
func NuevaVistaCar(car *models.Car, ctx *scene.Context) *CarView {

	sprite, err := render.LoadSprite(car.ModelPath)
	if err != nil {
		return nil
	}
	arribaSprite := sprite
	abajoSprite := sprite.Copy().Modify(mod.FlipY)
	izquierdaSprite := sprite.Copy().Modify(mod.Rotate(90))
	derechaSprite := sprite.Copy().Modify(mod.Rotate(-90))
	spriteSwitch := render.NewSwitch("arriba", map[string]render.Modifiable{
		"arriba":    arribaSprite,
		"abajo":     abajoSprite,
		"izquierda": izquierdaSprite,
		"derecha":   derechaSprite,
	})
	x, y := car.GetPosicion()
	spriteSwitch.SetPos(x, y)
	render.Draw(spriteSwitch, 3) 

	carView := &CarView{
		Car:     car,
		Sprite:  spriteSwitch,
		Context: ctx,
	}
	car.RegisterObserver(carView)
	return carView
}
func (cv *CarView) Update(data interface{}) {
	car := data.(*models.Car)
	x, y := car.GetPosicion()
	cv.Sprite.SetPos(x, y)

	direction := car.GetDirectionName()
	cv.Sprite.Set(direction)
}
