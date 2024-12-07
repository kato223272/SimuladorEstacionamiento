package models

import "sync"

// list autos
type GestionCar struct {
	mu   sync.Mutex
	Cars []*Car
}
func NuevoGestionCar() *GestionCar {
	return &GestionCar{
		Cars: []*Car{},
	}
}
func (cm *GestionCar) AddCarGestor(car *Car) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.Cars = append(cm.Cars, car)
}

func (cm *GestionCar) QuitarCar(car *Car) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for i, c := range cm.Cars {
		if c == car {
			cm.Cars = append(cm.Cars[:i], cm.Cars[i+1:]...)
			break
		}
	}
}
func (cm *GestionCar) GetCars() []*Car {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	carsCopy := make([]*Car, len(cm.Cars))
	copy(carsCopy, cm.Cars)
	return carsCopy
}
