package models

type Estacionamiento struct {
	Spots          []*EstacionamientoLugar
	ColaCars      *ColaCar
	AvailableSpots chan *EstacionamientoLugar
}


func NewEstacionamiento(spots []*EstacionamientoLugar) *Estacionamiento {
	availableSpots := make(chan *EstacionamientoLugar, len(spots))
	for _, spot := range spots {
		availableSpots <- spot
	}

	return &Estacionamiento{
		Spots:          spots,
		ColaCars:      NuevoCarCola(),
		AvailableSpots: availableSpots,
	}
}

func (p *Estacionamiento) GetDisponibleSpot() *EstacionamientoLugar {
	return <-p.AvailableSpots
}

func (p *Estacionamiento) LiberaSpot(spot *EstacionamientoLugar) {
	p.AvailableSpots <- spot
}
