package pueblo

import (
	"errors"
	"math"
)

const radioTierraMetros = 6371000.0

var (
	ErrLatitudInvalida  = errors.New("la latitud debe estar entre -90 y 90 grados")
	ErrLongitudInvalida = errors.New("la longitud debe estar entre -180 y 180 grados")
)

type Coordenada struct {
	Latitud  float64
	Longitud float64
}

func NuevaCoordenada(latitud, longitud float64) (Coordenada, error) {
	if math.IsNaN(latitud) || math.IsInf(latitud, 0) || latitud < -90 || latitud > 90 {
		return Coordenada{}, ErrLatitudInvalida
	}
	if math.IsNaN(longitud) || math.IsInf(longitud, 0) || longitud < -180 || longitud > 180 {
		return Coordenada{}, ErrLongitudInvalida
	}

	return Coordenada{
		Latitud:  latitud,
		Longitud: longitud,
	}, nil
}

func (c Coordenada) DistanciaA(otra Coordenada) float64 {
	lat1 := c.Latitud * math.Pi / 180
	lat2 := otra.Latitud * math.Pi / 180
	deltaLat := (otra.Latitud - c.Latitud) * math.Pi / 180
	deltaLon := (otra.Longitud - c.Longitud) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	anguloCentral := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return radioTierraMetros * anguloCentral
}
