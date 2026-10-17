package dominio

import (
	"errors"
	"fmt"
)

var (
	ErrCalleDuplicada   = errors.New("ya existe una calle con ese identificador")
	ErrCalleDesconocida = errors.New("la calle no está en el callejero")
)

type Callejero struct {
	calles map[string]*Calle
}

func NuevoCallejero() *Callejero {
	return &Callejero{calles: make(map[string]*Calle)}
}

func (c *Callejero) RegistrarCalle(id string, inicio, fin Coordenada, tipo string, arboles int) error {
	calle, err := nuevaCalle(id, inicio, fin, tipo, arboles)
	if err != nil {
		return err
	}
	if _, existe := c.calles[id]; existe {
		return fmt.Errorf("%w: %s", ErrCalleDuplicada, id)
	}
	if c.calles == nil {
		c.calles = make(map[string]*Calle)
	}
	c.calles[id] = calle
	return nil
}
