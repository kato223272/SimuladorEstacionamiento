package scenes

import (
	"estacionamiento/src/controllers"
	"estacionamiento/src/models"
	"estacionamiento/src/views"

	"github.com/oakmound/oak/v4/scene"
)

func NuevaEscenaEstacionamiento() *scene.Scene {
	return &scene.Scene{
		Start: func(ctx *scene.Context) {
			spots := []*models.EstacionamientoLugar{
				// Fila 1
				//1
				models.NewEstacionamientoLugar(380, 70, 410, 100, 1, 1),
				//2
				models.NewEstacionamientoLugar(425, 70, 455, 100, 1, 2),
				models.NewEstacionamientoLugar(470, 70, 500, 100, 1, 3),
				models.NewEstacionamientoLugar(515, 70, 545, 100, 1, 4),
				models.NewEstacionamientoLugar(560, 70, 590, 100, 1, 5),

				// Fila 2
				models.NewEstacionamientoLugar(380, 160, 410, 190, 2, 6),
				models.NewEstacionamientoLugar(425, 160, 455, 190, 2, 7),
				models.NewEstacionamientoLugar(470, 160, 500, 190, 2, 8),
				models.NewEstacionamientoLugar(515, 160, 545, 190, 2, 9),
				models.NewEstacionamientoLugar(560, 160, 590, 190, 2, 10),

				// Fila 3
				models.NewEstacionamientoLugar(380, 250, 410, 280, 3, 11),
				models.NewEstacionamientoLugar(425, 250, 455, 280, 3, 12),
				models.NewEstacionamientoLugar(470, 250, 500, 280, 3, 13),
				models.NewEstacionamientoLugar(515, 250, 545, 280, 3, 14),
				models.NewEstacionamientoLugar(560, 250, 590, 280, 3, 15),

				// Fila 4
				models.NewEstacionamientoLugar(380, 340, 410, 370, 4, 16),
				models.NewEstacionamientoLugar(425, 340, 455, 370, 4, 17),
				models.NewEstacionamientoLugar(470, 340, 500, 370, 4, 18),
				models.NewEstacionamientoLugar(515, 340, 545, 370, 4, 19),
				models.NewEstacionamientoLugar(560, 340, 590, 370, 4, 20),
			}

			Estacionamiento := models.NewEstacionamiento(spots)
			estacionamientoController := controllers.NuevoEstacionamientoController(Estacionamiento)
			EstacionamientoView := views.NewEstacionamientoView(Estacionamiento, ctx)
			estacionamientoController.View = EstacionamientoView

			estacionamientoController.InicioGeneracionCar(ctx)
		},
	}
}
