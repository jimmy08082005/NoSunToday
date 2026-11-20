package pueblo

import (
	"errors"
	"math"
)

const densidadSombraCompleta = 0.1

var ErrTramoLongitudCero = errors.New("no se puede calcular la frondosidad de un tramo con longitud cero")

type TipoCalle string

const (
	Ancha    TipoCalle = "ancha"
	Estrecha TipoCalle = "estrecha"
)

type Edificio struct {
	Plantas   int
	Ubicacion Coordenada
}

type Tramo struct {
	Tipo          TipoCalle
	NumeroArboles uint
	Inicio        Coordenada
	Fin           Coordenada
	Edificios     []Edificio
}

func (t Tramo) Longitud() float64 {
	return t.Inicio.DistanciaA(t.Fin)
}

func (t Tramo) Frondosidad() (float64, error) {
	longitud := t.Longitud()
	if longitud <= 0 {
		return 0, ErrTramoLongitudCero
	}

	densidad := float64(t.NumeroArboles) / longitud
	indice := densidad / densidadSombraCompleta

	return math.Min(1.0, indice), nil
}