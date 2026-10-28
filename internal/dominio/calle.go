package dominio

import "errors"
import "maps"

var (
	ErrArbolesNegativos = errors.New("el número de árboles no puede ser negativo")
)

type Calle struct {
	tipo      string
	arboles   int
	edificios map[Coordenada]Edificio
}

func NuevaCalle(tipo string, arboles int, edificios map[Coordenada]Edificio) (*Calle, error) {
	if arboles < 0 {
		return nil, ErrArbolesNegativos
	}
	return &Calle{
		tipo:      tipo,
		arboles:   arboles,
		edificios: maps.Clone(edificios),
	}, nil
}
