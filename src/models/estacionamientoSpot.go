package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)
type DireccionLugarEstacionamiento struct {
	Direction string 
	Point     float64
}
func NewDireccionLugarEstacionamiento(direction string, point float64) *DireccionLugarEstacionamiento {
	return &DireccionLugarEstacionamiento{
		Direction: direction,
		Point:     point,
	}
}
type EstacionamientoLugar struct {
	Area                 *floatgeom.Rect2     //delimita el lugar
	DirectionsForEstacionamiento []*DireccionLugarEstacionamiento 
	SalirEstacionamiento []*DireccionLugarEstacionamiento
	Number               int
	IsAvailable          bool
}

func NewEstacionamientoLugar(x, y, x2, y2 float64, row, number int) *EstacionamientoLugar {
	directionsForEstacionamiento := GetDireccionEstacionamiento(x, y, row)
	SalirEstacionamiento := GetSalirEstacionamiento()
	area := floatgeom.NewRect2(x, y, x2, y2)

	return &EstacionamientoLugar{
		Area:                 &area,
		DirectionsForEstacionamiento: directionsForEstacionamiento,
		SalirEstacionamiento: SalirEstacionamiento,
		Number:               number,
		IsAvailable:          true,
	}
}

// estacionar de la fila
func GetDireccionEstacionamiento(x, y float64, row int) []*DireccionLugarEstacionamiento {
	var directions []*DireccionLugarEstacionamiento

	switch row {
	case 1:
		directions = append(directions, NewDireccionLugarEstacionamiento("abajo", 45))
	case 2:
		directions = append(directions, NewDireccionLugarEstacionamiento("abajo", 135))
	case 3:
		directions = append(directions, NewDireccionLugarEstacionamiento("abajo", 225))
	case 4:
		directions = append(directions, NewDireccionLugarEstacionamiento("abajo", 315))
	}

	directions = append(directions, NewDireccionLugarEstacionamiento("derecha", x+5))
	directions = append(directions, NewDireccionLugarEstacionamiento("abajo", y+5))

	return directions
}

func GetSalirEstacionamiento() []*DireccionLugarEstacionamiento {
	var directions []*DireccionLugarEstacionamiento

	directions = append(directions, NewDireccionLugarEstacionamiento("derecha", 600))
	directions = append(directions, NewDireccionLugarEstacionamiento("arriba", 15))
	directions = append(directions, NewDireccionLugarEstacionamiento("izquierda", 250))
	return directions
}
func (p *EstacionamientoLugar) GetArea() *floatgeom.Rect2 {
	return p.Area
}
func (p *EstacionamientoLugar) GetNumeroLugar() int {
	return p.Number
}
func (p *EstacionamientoLugar) GetDireccionEstacionamiento() []*DireccionLugarEstacionamiento {
	return p.DirectionsForEstacionamiento
}
func (p *EstacionamientoLugar) GetSalirEstacionamiento() []*DireccionLugarEstacionamiento {
	return p.SalirEstacionamiento
}
func (p *EstacionamientoLugar) GetLugarDisponible() bool {
	return p.IsAvailable
}
func (p *EstacionamientoLugar) SetLugarDisponible(isAvailable bool) {
	p.IsAvailable = isAvailable
}
//obtener coordenadas area
func (p *EstacionamientoLugar) GetMinimaX() float64 {
	return p.Area.Min.X()
}
func (p *EstacionamientoLugar) GetMinimaY() float64 {
	return p.Area.Min.Y()
}
func (p *EstacionamientoLugar) GetMaximaX() float64 {
	return p.Area.Max.X()
}
func (p *EstacionamientoLugar) GetMinimaY2() float64 {
	return p.Area.Max.Y()
}
