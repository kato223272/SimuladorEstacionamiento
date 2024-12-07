package controllers

import (
	"estacionamiento/src/models"
	"estacionamiento/src/views"
	
	"math/rand"
	"time"
)
type CarController struct {
	Car        *models.Car
	Estacionamiento    *models.Estacionamiento
	CarView    *views.CarView
	GestionCar *models.GestionCar
	signalChan   chan struct{}
	abandonarChan   chan struct{} //priori a autos que salen
}
func NuevoCarController(car *models.Car, Estacionamiento *models.Estacionamiento, GestionCar *models.GestionCar, signalChan chan struct{}, abandonarChan chan struct{}) *CarController {
	return &CarController{
		Car:        car,
		Estacionamiento:    Estacionamiento,
		GestionCar: GestionCar,
		signalChan:   signalChan,
		abandonarChan:   abandonarChan,
	}
}
func (cc *CarController) Start() {
	cc.GestionCar.AddCarGestor(cc.Car)

	cc.cola()

	spot := cc.Estacionamiento.GetDisponibleSpot()

	cc.Park(spot)

	time.Sleep(time.Second * time.Duration(rand.Intn(15)+20))

	cc.LeaveSpot()
	cc.Estacionamiento.LiberaSpot(spot)

	cc.Leave(spot)
	
	<-cc.abandonarChan
	cc.ExitDoor()
	cc.abandonarChan <- struct{}{}

	cc.AbandonarEstacionamiento()
	cc.GestionCar.QuitarCar(cc.Car)
}

//mov del carro
func (cc *CarController) cola() {
	cc.Estacionamiento.ColaCars.AniadeCola(cc.Car)

	minY := 45.0    
	spacing := 50.0 //dist

	for cc.Car.Y > minY {
		carAhead := cc.Estacionamiento.ColaCars.GetCarAlfrente(cc.Car)
		canMove := true

		if carAhead != nil {
			_, aheadY := carAhead.GetPosicion()
			ccY := cc.Car.Y

			if ccY-aheadY < spacing {
				canMove = false
			}
		}

		if canMove {
			cc.Car.SetDireccion(0, -1)
			cc.Car.Move(0, -1)
		} else {
			cc.Car.SetDireccion(0, -1) //mantiene
		}
		time.Sleep(10 * time.Millisecond)
	}

	<-cc.signalChan
	defer func() { cc.signalChan <- struct{}{} }()//avan

	<-cc.abandonarChan
	defer func() { cc.abandonarChan <- struct{}{} }()

	cc.JoinDoor()

	cc.Estacionamiento.ColaCars.QuitarCar(cc.Car)
}

func (cc *CarController) JoinDoor() {
	minDistance := 50.0                                    
	for cc.Car.X < 355 {
		canMove := true

		for _, otherCar := range cc.GestionCar.GetCars() {
			if otherCar != cc.Car {
				otherX, otherY := otherCar.GetPosicion()
				if cc.Car.Y == otherY && cc.Car.X < otherX && otherX-cc.Car.X < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			cc.Car.SetDireccion(1, 0) 
			cc.Car.Move(1, 0)
		} else {
			cc.Car.SetDireccion(1, 0) 
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) Park(spot *models.EstacionamientoLugar) {
	for _, direction := range spot.GetDireccionEstacionamiento() {
		cc.move(direction)
	}
}

func (cc *CarController) LeaveSpot() {
	cc.Car.SetDireccion(0, -1)                               
	cc.Car.Move(0, -30)
}

func (cc *CarController) Leave(spot *models.EstacionamientoLugar) {
	for _, direction := range spot.GetSalirEstacionamiento() {
		cc.move(direction)
	}
}

func (cc *CarController) ExitDoor() {
	minDistance := 50.0                       
	for cc.Car.X > 300 {
		canMove := true
		for _, otherCar := range cc.GestionCar.GetCars() {
			if otherCar != cc.Car {
				otherX, otherY := otherCar.GetPosicion()
				if cc.Car.Y == otherY && cc.Car.X > otherX && cc.Car.X-otherX < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			cc.Car.SetDireccion(-1, 0) 
			cc.Car.Move(-1, 0)
		} else {
			cc.Car.SetDireccion(-1, 0)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) AbandonarEstacionamiento() {
	cc.Car.SetDireccion(-1, 0)   
	for cc.Car.X > -5 {
		cc.Car.Move(-1, 0)
		time.Sleep(5 * time.Millisecond)
	}

	cc.GestionCar.QuitarCar(cc.Car)
}

func (cc *CarController) move(direction *models.DireccionLugarEstacionamiento) {
	minDistance := 50.0 
	for {
		var canMove bool = true
		var dx, dy float64

		switch direction.Direction {
		case "izquierda":
			if cc.Car.X <= direction.Point {
				return
			}
			dx, dy = -1, 0
			for _, otherCar := range cc.GestionCar.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosicion()
					if cc.Car.Y == otherY && cc.Car.X > otherX && cc.Car.X-otherX < minDistance {
						canMove = false
						break
					}
				}
			}
		case "derecha":
			if cc.Car.X >= direction.Point {
				return
			}
			dx, dy = 1, 0
			for _, otherCar := range cc.GestionCar.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosicion()
					if cc.Car.Y == otherY && cc.Car.X < otherX && otherX-cc.Car.X < minDistance {
						canMove = false
						break
					}
				}
			}
		case "arriba":
			if cc.Car.Y <= direction.Point {
				return
			}
			dx, dy = 0, -1
			for _, otherCar := range cc.GestionCar.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosicion()
					if cc.Car.X == otherX && cc.Car.Y > otherY && cc.Car.Y-otherY < minDistance {
						canMove = false
						break
					}
				}
			}
		case "abajo":
			if cc.Car.Y >= direction.Point {
				return
			}
			dx, dy = 0, 1
			for _, otherCar := range cc.GestionCar.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosicion()
					if cc.Car.X == otherX && cc.Car.Y < otherY && otherY-cc.Car.Y < minDistance {
						canMove = false
						break
					}
				}
			}
		}

		if canMove {
			cc.Car.SetDireccion(dx, dy) //dire actual
			cc.Car.Move(dx, dy)
		} else {
			cc.Car.SetDireccion(dx, dy)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
