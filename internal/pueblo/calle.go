package pueblo

import (
	"errors"
	"math"
)

var ErrCalleSinLongitud = errors.New("no se puede calcular la frondosidad de una calle sin tramos o con longitud cero")

type Calle struct {
	Tramos []Tramo
}

func (c Calle) Longitud() float64 {
	var total float64
	for _, tramo := range c.Tramos {
		total += tramo.Longitud()
	}
	return total
}

func (c Calle) Frondosidad() (float64, error) {
	var longitudTotal float64
	var arbolesTotales uint

	for _, tramo := range c.Tramos {
		longitudTotal += tramo.Longitud()
		arbolesTotales += tramo.NumeroArboles
	}

	if longitudTotal <= 0 {
		return 0, ErrCalleSinLongitud
	}

	densidad := float64(arbolesTotales) / longitudTotal
	indice := densidad / densidadSombraCompleta

	return math.Min(1.0, indice), nil
}
