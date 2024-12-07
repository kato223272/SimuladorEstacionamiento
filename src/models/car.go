package models

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	mu        sync.Mutex
	X         float64
	Y         float64
	DX        float64 
	DY        float64 
	ModelPath string
	observers []Observer
}

func NewCar() *Car {
	modelPaths := []string{
		"assets/img/cafe.png",
		"assets/img/verde.png",
		"assets/img/naranja.png",
		"assets/img/blanco.png",
	}
	rand.Seed(time.Now().UnixNano())
	modelPath := modelPaths[rand.Intn(len(modelPaths))]

	return &Car{
		X:         300, 
		Y:         400, 
		DX:        0,
		DY:        -1, 
		ModelPath: modelPath,
		observers: []Observer{},
	}
}

func (c *Car) RegisterObserver(o Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.observers = append(c.observers, o)
}

func (c *Car) RemoveObserver(o Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, observer := range c.observers {
		if observer == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *Car) NotifyObservers() {
	c.mu.Lock()
	observers := make([]Observer, len(c.observers))
	copy(observers, c.observers)
	c.mu.Unlock()

	for _, observer := range observers {
		observer.Update(c)
	}
}

func (c *Car) Move(dx, dy float64) {
	c.mu.Lock()
	c.X += dx
	c.Y += dy
	c.mu.Unlock()
	c.NotifyObservers()
}

func (c *Car) SetDireccion(dx, dy float64) {
	c.mu.Lock()
	c.DX = dx
	c.DY = dy
	c.mu.Unlock()
	c.NotifyObservers()
}

func (c *Car) GetPosicion() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.X, c.Y
}

func (c *Car) GetDirection() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.DX, c.DY
}

func (c *Car) GetDirectionName() string {
	c.mu.Lock()
	dx := c.DX
	dy := c.DY
	c.mu.Unlock()

	if dx == 0 && dy == 0 {
		return "arriba" 
	}

	angle := math.Atan2(dy, dx) * (180 / math.Pi)

	if angle >= -45 && angle <= 45 {
		return "derecha"
	} else if angle > 45 && angle < 135 {
		return "abajo" 
	} else if angle >= 135 || angle <= -135 {
		return "izquierda"
	} else {
		return "arriba" 
	}
}

func (c *Car) SetX(x float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.X = x
	c.NotifyObservers()
}

func (c *Car) SetY(y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Y = y
	c.NotifyObservers()
}
