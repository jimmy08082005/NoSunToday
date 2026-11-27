package pueblo

import (
	"errors"
	"math"
)

const densidadSombraCompleta = 0.1

var (
	ErrTramoLongitudCero  = errors.New("no se puede calcular la frondosidad de un tramo con longitud cero")
	ErrPesoSombraInvalido = errors.New("el peso de la sombra debe estar en el intervalo [0, 1)")
)

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

const factorSombra = 0.65

func (t Tramo) Coste() (float64, error) {
	frondosidad, err := t.Frondosidad()
	if err != nil {
		return 0, err
	}

	return t.Longitud() * (1.0 - factorSombra*frondosidad), nil
}