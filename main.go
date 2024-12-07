package main

import (
	"estacionamiento/interface"

	"github.com/oakmound/oak/v4"
)

func main() {
	scene := scenes.NuevaEscenaEstacionamiento()
	oak.AddScene("EstacionamientoEscena", *scene)
	err := oak.Init("EstacionamientoEscena", func(c oak.Config) (oak.Config, error) {
		c.Screen.Width = 900
		c.Screen.Height = 400
		c.Assets.ImagePath = "assets/img"
		return c, nil
	})
	if err != nil {
		panic(err)
	}
}
