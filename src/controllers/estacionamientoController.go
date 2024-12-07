package controllers

import (
	"estacionamiento/src/models"
	"estacionamiento/src/views"
	"math/rand"
	"time"

	"github.com/oakmound/oak/v4/scene"
)

type estacionamientoController struct {
	Estacionamiento    *models.Estacionamiento
	View       *views.EstacionamientoView
	signalChan   chan struct{}
	abandonarChan   chan struct{}
	GestionCar *models.GestionCar
}


func NuevoEstacionamientoController(Estacionamiento *models.Estacionamiento) *estacionamientoController {
	doorChan := make(chan struct{}, 1)
	doorChan <- struct{}{} 

	pathChan := make(chan struct{}, 1)
	pathChan <- struct{}{} 

	return &estacionamientoController{
		Estacionamiento:    Estacionamiento,
		signalChan:   doorChan,
		abandonarChan:   pathChan,
		GestionCar: models.NuevoGestionCar(),
	}
}

func (pc *estacionamientoController) InicioGeneracionCar(ctx *scene.Context) {
	const maxCars = 100
	go func() {
		for i := 0; i < maxCars; i++ {
			car := models.NewCar()
			carController := NuevoCarController(car, pc.Estacionamiento, pc.GestionCar, pc.signalChan, pc.abandonarChan)
			carView := views.NuevaVistaCar(car, ctx)
			carController.CarView = carView
			go carController.Start()
			time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
		}
	}()
}
