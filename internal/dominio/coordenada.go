package dominio

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrLatitudInvalida  = errors.New("latitud inválida: debe estar entre -90 y 90")
	ErrLongitudInvalida = errors.New("longitud inválida: debe estar entre -180 y 180")
)

type Latitud float64
type Longitud float64

type Coordenada struct {
	latitud  Latitud
	longitud Longitud
}

func NuevaCoordenada(lat Latitud, lon Longitud) (Coordenada, error) {
	if math.IsNaN(float64(lat)) || lat < -90.0 || lat > 90.0 {
		return Coordenada{}, fmt.Errorf("%w (valor recibido: %v)", ErrLatitudInvalida, lat)
	}
	if math.IsNaN(float64(lon)) || lon < -180.0 || lon > 180.0 {
		return Coordenada{}, fmt.Errorf("%w (valor recibido: %v)", ErrLongitudInvalida, lon)
	}
	return Coordenada{latitud: lat, longitud: lon}, nil
}
