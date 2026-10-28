package dominio

import (
	"errors"
	"fmt"
)

var ErrPlantasInvalidas = errors.New("un edificio debe tener al menos una planta")

type Edificio struct {
	ubicacion Coordenada
	plantas   int
}

func NuevoEdificio(ubicacion Coordenada, plantas int) (Edificio, error) {
	if plantas < 1 {
		return Edificio{}, fmt.Errorf("%w (valor recibido: %d)", ErrPlantasInvalidas, plantas)
	}
	return Edificio{ubicacion: ubicacion, plantas: plantas}, nil
}
