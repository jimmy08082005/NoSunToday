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

func nuevoEdificio(ubicacion Coordenada, plantas int) (Edificio, error) {
	if plantas < 1 {
		return Edificio{}, fmt.Errorf("%w (valor recibido: %d)", ErrPlantasInvalidas, plantas)
	}
	return Edificio{ubicacion: ubicacion, plantas: plantas}, nil
}

func (c *Callejero) RegistrarEdificio(idCalle string, ubicacion Coordenada, plantas int) error {
	calle, existe := c.calles[idCalle]
	if !existe {
		return fmt.Errorf("%w: %s", ErrCalleDesconocida, idCalle)
	}
	edificio, err := nuevoEdificio(ubicacion, plantas)
	if err != nil {
		return err
	}
	calle.edificios = append(calle.edificios, edificio)
	return nil
}
