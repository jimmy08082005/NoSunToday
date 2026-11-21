package pueblo

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrCalleNoEncontrada = errors.New("la calle indicada no existe o no tiene tramos en el callejero")
	ErrRutaNoEncontrada   = errors.New("no existe un camino conectado entre la calle de origen y la de destino")
)

func generarGrafo(callejero Callejero) map[Coordenada]map[NombreCalle][]Tramo {
	grafo := make(map[Coordenada]map[NombreCalle][]Tramo)

	for nombre, calle := range callejero.Calles {
		for _, tramo := range calle.Tramos {
			if grafo[tramo.Inicio] == nil {
				grafo[tramo.Inicio] = make(map[NombreCalle][]Tramo)
			}
			grafo[tramo.Inicio][nombre] = append(grafo[tramo.Inicio][nombre], tramo)

			if grafo[tramo.Fin] == nil {
				grafo[tramo.Fin] = make(map[NombreCalle][]Tramo)
			}
			grafo[tramo.Fin][nombre] = append(grafo[tramo.Fin][nombre], tramo)
		}
	}

	return grafo
}

func CalcularRuta(callejero Callejero, origen, destino NombreCalle) (Callejero, error) {
	calleOrigen, okOrigen := callejero.Calles[origen]
	if !okOrigen || len(calleOrigen.Tramos) == 0 {
		return Callejero{}, fmt.Errorf("%w: %s", ErrCalleNoEncontrada, origen)
	}

	calleDestino, okDestino := callejero.Calles[destino]
	if !okDestino || len(calleDestino.Tramos) == 0 {
		return Callejero{}, fmt.Errorf("%w: %s", ErrCalleNoEncontrada, destino)
	}

	if origen == destino {
		return Callejero{
			Calles: map[NombreCalle]Calle{
				origen: calleOrigen,
			},
		}, nil
	}

	grafo := generarGrafo(callejero)

	nodosDestino := make(map[Coordenada]bool)
	for _, tramo := range calleDestino.Tramos {
		nodosDestino[tramo.Inicio] = true
		nodosDestino[tramo.Fin] = true
	}

	costes := make(map[Coordenada]float64)
	previos := make(map[Coordenada]pasoPrevio)
	cola := &colaPrioridad{}
	heap.Init(cola)

	for _, tramo := range calleOrigen.Tramos {
		for _, punto := range []Coordenada{tramo.Inicio, tramo.Fin} {
			if _, yaIniciado := costes[punto]; !yaIniciado {
				costes[punto] = 0
				heap.Push(cola, nodoCola{coord: punto, coste: 0})
			}
		}
	}

	for cola.Len() > 0 {
		actual := heap.Pop(cola).(nodoCola)

		if actual.coste > costes[actual.coord] {
			continue
		}

		if nodosDestino[actual.coord] {
			return reconstruirCallejeroRuta(callejero, origen, destino, actual.coord, previos), nil
		}

		for nombreCalle, tramos := range grafo[actual.coord] {
			for _, tramo := range tramos {
				siguiente := tramo.Fin
				if actual.coord == tramo.Fin {
					siguiente = tramo.Inicio
				}

				costeTramo, err := tramo.Coste()
				if err != nil {
					return Callejero{}, err
				}

				nuevoCoste := actual.coste + costeTramo
				costeAnterior, visitado := costes[siguiente]

				if !visitado || nuevoCoste < costeAnterior {
					costes[siguiente] = nuevoCoste
					previos[siguiente] = pasoPrevio{
						desde: actual.coord,
						calle: nombreCalle,
						tramo: tramo,
					}
					heap.Push(cola, nodoCola{coord: siguiente, coste: nuevoCoste})
				}
			}
		}
	}

	return Callejero{}, ErrRutaNoEncontrada
}

type pasoPrevio struct {
	desde Coordenada
	calle NombreCalle
	tramo Tramo
}

func reconstruirCallejeroRuta(callejero Callejero, origen, destino NombreCalle, meta Coordenada, previos map[Coordenada]pasoPrevio) Callejero {
	var pasos []pasoPrevio
	actual := meta

	for {
		paso, existe := previos[actual]
		if !existe {
			break
		}
		pasos = append(pasos, paso)
		actual = paso.desde
	}

	callesRuta := make(map[NombreCalle]Calle)

	for i := len(pasos) - 1; i >= 0; i-- {
		paso := pasos[i]
		calle := callesRuta[paso.calle]
		calle.Tramos = append(calle.Tramos, paso.tramo)
		callesRuta[paso.calle] = calle
	}

	if _, ok := callesRuta[origen]; !ok {
		callesRuta[origen] = callejero.Calles[origen]
	}
	if _, ok := callesRuta[destino]; !ok {
		callesRuta[destino] = callejero.Calles[destino]
	}

	return Callejero{Calles: callesRuta}
}

type nodoCola struct {
	coord Coordenada
	coste float64
}

type colaPrioridad []nodoCola

func (c colaPrioridad) Len() int           { return len(c) }
func (c colaPrioridad) Less(i, j int) bool { return c[i].coste < c[j].coste }
func (c colaPrioridad) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c *colaPrioridad) Push(x any)        { *c = append(*c, x.(nodoCola)) }
func (c *colaPrioridad) Pop() any {
	antiguo := *c
	n := len(antiguo)
	item := antiguo[n-1]
	*c = antiguo[:n-1]
	return item
}