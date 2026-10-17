package dominio

import "errors"

var (
	ErrIDVacio          = errors.New("la calle debe tener un identificador")
	ErrArbolesNegativos = errors.New("el número de árboles no puede ser negativo")
)

type Calle struct {
	id      string
	inicio  Coordenada
	fin     Coordenada
	tipo    string
	arboles int
}

func nuevaCalle(id string, inicio, fin Coordenada, tipo string, arboles int) (*Calle, error) {
	if id == "" {
		return nil, ErrIDVacio
	}
	if arboles < 0 {
		return nil, ErrArbolesNegativos
	}

	return &Calle{
		id:      id,
		inicio:  inicio,
		fin:     fin,
		tipo:    tipo,
		arboles: arboles,
	}, nil
}
