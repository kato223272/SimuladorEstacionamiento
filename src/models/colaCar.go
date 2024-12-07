package models

import (
	"sync"
)

type ColaCar struct {
	mu   sync.Mutex
	cars []*Car 
}

func NuevoCarCola() *ColaCar {
	return &ColaCar{
		cars: []*Car{},
	}
}

func (cq *ColaCar) AniadeCola(car *Car) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.cars = append(cq.cars, car)
}

func (cq *ColaCar) EliminaDeCola() *Car {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	if len(cq.cars) == 0 {
		return nil
	}
	car := cq.cars[0]
	cq.cars = cq.cars[1:]
	return car
}

func (cq *ColaCar) GetposicionCola(car *Car) int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	for i, c := range cq.cars {
		if c == car {
			return i
		}
	}
	return -1
}

//Delante en cola
func (cq *ColaCar) GetCarAlfrente(car *Car) *Car {
	position := cq.GetposicionCola(car)
	if position > 0 {
		return cq.cars[position-1]
	}
	return nil
}

func (cq *ColaCar) QuitarCar(car *Car) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	for i, c := range cq.cars {
		if c == car {
			cq.cars = append(cq.cars[:i], cq.cars[i+1:]...)
			break
		}
	}
}
