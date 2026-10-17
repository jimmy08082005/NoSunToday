package dominio

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrRutaVacia       = errors.New("una ruta debe contener al menos una calle")
	ErrRutaDiscontinua = errors.New("las calles de la ruta no están conectadas")
)

type Ruta struct {
	calles []string
}

func (r Ruta) Igual(otra Ruta) bool {
	return slices.Equal(r.calles, otra.calles)
}

func (c *Callejero) NuevaRuta(ids []string) (Ruta, error) {
	if len(ids) == 0 {
		return Ruta{}, ErrRutaVacia
	}

	calles := make([]*Calle, 0, len(ids))
	for _, id := range ids {
		calle, existe := c.calles[id]
		if !existe {
			return Ruta{}, fmt.Errorf("%w: %s", ErrCalleDesconocida, id)
		}
		calles = append(calles, calle)
	}

	if i := primeraDesconectada(calles); i >= 0 {
		return Ruta{}, fmt.Errorf("%w: la calle %s no conecta con las anteriores", ErrRutaDiscontinua, calles[i].id)
	}

	return Ruta{calles: append([]string(nil), ids...)}, nil
}

func primeraDesconectada(calles []*Calle) int {
	salidas := []Coordenada{calles[0].inicio, calles[0].fin}
	for i := 1; i < len(calles); i++ {
		var nuevas []Coordenada
		for _, p := range salidas {
			if p == calles[i].inicio {
				nuevas = append(nuevas, calles[i].fin)
			}
			if p == calles[i].fin {
				nuevas = append(nuevas, calles[i].inicio)
			}
		}
		if len(nuevas) == 0 {
			return i
		}
		salidas = nuevas
	}
	return -1
}
