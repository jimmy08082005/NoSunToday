package pueblo

import (
	"container/heap"
	"errors"
	"strings"
	"testing"
)

func crearCallejeroDePrueba() Callejero {
	puntoA := Coordenada{Latitud: 37.2220, Longitud: -3.6880}
	puntoB := Coordenada{Latitud: 37.2230, Longitud: -3.6880}
	puntoD := Coordenada{Latitud: 37.2240, Longitud: -3.6880}
	puntoE := Coordenada{Latitud: 37.2250, Longitud: -3.6880}
	puntoC1 := Coordenada{Latitud: 37.2235, Longitud: -3.6874}
	puntoC2 := Coordenada{Latitud: 37.2235, Longitud: -3.6860}

	isla1 := Coordenada{Latitud: 37.2300, Longitud: -3.6900}
	isla2 := Coordenada{Latitud: 37.2310, Longitud: -3.6900}

	return Callejero{
		Calles: map[NombreCalle]Calle{
			"Calle Origen": {
				Tramos: []Tramo{{Inicio: puntoA, Fin: puntoB, NumeroArboles: 0}},
			},
			"Calle Destino": {
				Tramos: []Tramo{{Inicio: puntoD, Fin: puntoE, NumeroArboles: 0}},
			},
			"Calle Soleada": {
				Tramos: []Tramo{{Inicio: puntoB, Fin: puntoD, NumeroArboles: 0}},
			},
			"Calle Arbolada": {
				Tramos: []Tramo{
					{Inicio: puntoB, Fin: puntoC1, NumeroArboles: 10},
					{Inicio: puntoC1, Fin: puntoD, NumeroArboles: 10},
				},
			},
			"Calle Rodeo Largo": {
				Tramos: []Tramo{
					{Inicio: puntoB, Fin: puntoC2, NumeroArboles: 25},
					{Inicio: puntoC2, Fin: puntoD, NumeroArboles: 25},
				},
			},
			"Calle Aislada": {
				Tramos: []Tramo{{Inicio: isla1, Fin: isla2, NumeroArboles: 5}},
			},
		},
	}
}

func rutaBasePorDistancia(callejero Callejero, origen, destino NombreCalle) (Callejero, error) {
	calleOrigen := callejero.Calles[origen]
	calleDestino := callejero.Calles[destino]
	grafo := generarGrafo(callejero)

	nodosDestino := make(map[Coordenada]bool)
	for _, tramo := range calleDestino.Tramos {
		nodosDestino[tramo.Inicio] = true
		nodosDestino[tramo.Fin] = true
	}

	distancias := make(map[Coordenada]float64)
	previos := make(map[Coordenada]pasoPrevio)
	cola := &colaPrioridad{}
	heap.Init(cola)

	for _, tramo := range calleOrigen.Tramos {
		for _, punto := range []Coordenada{tramo.Inicio, tramo.Fin} {
			if _, ok := distancias[punto]; !ok {
				distancias[punto] = 0
				heap.Push(cola, nodoCola{coord: punto, coste: 0})
			}
		}
	}

	for cola.Len() > 0 {
		actual := heap.Pop(cola).(nodoCola)
		if actual.coste > distancias[actual.coord] {
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
				nuevaDist := actual.coste + tramo.Longitud()
				if ant, visitado := distancias[siguiente]; !visitado || nuevaDist < ant {
					distancias[siguiente] = nuevaDist
					previos[siguiente] = pasoPrevio{desde: actual.coord, calle: nombreCalle, tramo: tramo}
					heap.Push(cola, nodoCola{coord: siguiente, coste: nuevaDist})
				}
			}
		}
	}

	return Callejero{}, ErrRutaNoEncontrada
}

func exposicionSolMetros(ruta Callejero) float64 {
	var metrosAlSol float64
	for _, calle := range ruta.Calles {
		for _, tramo := range calle.Tramos {
			frondosidad, err := tramo.Frondosidad()
			if err != nil {
				continue
			}
			metrosAlSol += tramo.Longitud() * (1.0 - frondosidad)
		}
	}
	return metrosAlSol
}

func distanciaTotalRuta(ruta Callejero) float64 {
	var total float64
	for _, calle := range ruta.Calles {
		for _, tramo := range calle.Tramos {
			total += tramo.Longitud()
		}
	}
	return total
}

func TestMejoraSobreCasoBase(t *testing.T) {
	t.Run("reduce la exposicion al sol frente al caso base sin penalizar en exceso la distancia", func(t *testing.T) {
		callejero := crearCallejeroDePrueba()
		delete(callejero.Calles, "Calle Rodeo Largo")

		rutaBase, err := rutaBasePorDistancia(callejero, "Calle Origen", "Calle Destino")
		if err != nil {
			t.Fatalf("error calculando ruta base: %v", err)
		}

		rutaOptima, err := CalcularRuta(callejero, "Calle Origen", "Calle Destino")
		if err != nil {
			t.Fatalf("error calculando ruta optima: %v", err)
		}

		solBase := exposicionSolMetros(rutaBase)
		solOptima := exposicionSolMetros(rutaOptima)

		if solOptima >= solBase {
			t.Errorf("la ruta optima no mejora la proteccion solar del caso base: optima=%.2f m al sol, base=%.2f m al sol", solOptima, solBase)
		}

		distBase := distanciaTotalRuta(rutaBase)
		distOptima := distanciaTotalRuta(rutaOptima)

		if distOptima > distBase*1.5 {
			t.Errorf("el rodeo es demasiado largo respecto al caso base: optima=%.2f m, base=%.2f m", distOptima, distBase)
		}
	})

	t.Run("no empeora el caso base cuando los rodeos con sombra son desproporcionados", func(t *testing.T) {
		callejero := crearCallejeroDePrueba()
		delete(callejero.Calles, "Calle Arbolada")

		rutaBase, err := rutaBasePorDistancia(callejero, "Calle Origen", "Calle Destino")
		if err != nil {
			t.Fatalf("error calculando ruta base: %v", err)
		}

		rutaOptima, err := CalcularRuta(callejero, "Calle Origen", "Calle Destino")
		if err != nil {
			t.Fatalf("error calculando ruta optima: %v", err)
		}

		distBase := distanciaTotalRuta(rutaBase)
		distOptima := distanciaTotalRuta(rutaOptima)

		if distOptima > distBase {
			t.Errorf("ante un rodeo excesivo deberia mantener la distancia del caso base: optima=%.2f m, base=%.2f m", distOptima, distBase)
		}
	})

	t.Run("mejora el caso base cargando datos estructurados de OpenStreetMap", func(t *testing.T) {
		jsonOSM := `{
			"elements": [
				{"type": "node", "id": 1, "lat": 37.2220, "lon": -3.6880},
				{"type": "node", "id": 2, "lat": 37.2230, "lon": -3.6880},
				{"type": "node", "id": 3, "lat": 37.2240, "lon": -3.6880},
				{"type": "node", "id": 4, "lat": 37.2250, "lon": -3.6880},
				{"type": "node", "id": 5, "lat": 37.2235, "lon": -3.6874},
				{"type": "node", "id": 101, "lat": 37.22325, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 102, "lat": 37.22325, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 103, "lat": 37.22325, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 104, "lat": 37.22325, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 105, "lat": 37.22325, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 106, "lat": 37.22375, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 107, "lat": 37.22375, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 108, "lat": 37.22375, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 109, "lat": 37.22375, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "node", "id": 110, "lat": 37.22375, "lon": -3.6877, "tags": {"natural": "tree"}},
				{"type": "way", "id": 10, "nodes": [1, 2], "tags": {"highway": "residential", "name": "Calle Inicio"}},
				{"type": "way", "id": 11, "nodes": [2, 3], "tags": {"highway": "residential", "name": "Via Directa"}},
				{"type": "way", "id": 12, "nodes": [2, 5, 3], "tags": {"highway": "residential", "name": "Paseo Arbolado"}},
				{"type": "way", "id": 13, "nodes": [3, 4], "tags": {"highway": "residential", "name": "Calle Meta"}}
			]
		}`

		callejero, err := CargarCallejeroDesdeOSM(strings.NewReader(jsonOSM))
		if err != nil {
			t.Fatalf("error cargando datos OSM: %v", err)
		}

		rutaBase, err := rutaBasePorDistancia(callejero, "Calle Inicio", "Calle Meta")
		if err != nil {
			t.Fatalf("error calculando ruta base: %v", err)
		}

		rutaOptima, err := CalcularRuta(callejero, "Calle Inicio", "Calle Meta")
		if err != nil {
			t.Fatalf("error calculando ruta optima: %v", err)
		}

		if exposicionSolMetros(rutaOptima) >= exposicionSolMetros(rutaBase) {
			t.Errorf("la ruta optima sobre datos OSM deberia reducir los metros al sol respecto al caso base")
		}
	})
}

func TestCalcularRutaErrores(t *testing.T) {
	callejero := crearCallejeroDePrueba()

	casos := []struct {
		nombre        string
		callejero     Callejero
		origen        NombreCalle
		destino       NombreCalle
		errorEsperado error
	}{
		{
			nombre:        "falla si la calle de origen no existe",
			callejero:     callejero,
			origen:        "Calle Inventada",
			destino:       "Calle Destino",
			errorEsperado: ErrCalleNoEncontrada,
		},
		{
			nombre:        "falla si la calle de destino no existe",
			callejero:     callejero,
			origen:        "Calle Origen",
			destino:       "Calle Inventada",
			errorEsperado: ErrCalleNoEncontrada,
		},
		{
			nombre:        "falla si no existe conexion entre origen y destino",
			callejero:     callejero,
			origen:        "Calle Origen",
			destino:       "Calle Aislada",
			errorEsperado: ErrRutaNoEncontrada,
		},
		{
			nombre: "falla si un tramo tiene longitud cero",
			callejero: Callejero{
				Calles: map[NombreCalle]Calle{
					"Origen": {Tramos: []Tramo{{
						Inicio: Coordenada{Latitud: 37.2220, Longitud: -3.6880},
						Fin:    Coordenada{Latitud: 37.2230, Longitud: -3.6880},
					}}},
					"Tramo Nulo": {Tramos: []Tramo{{
						Inicio: Coordenada{Latitud: 37.2230, Longitud: -3.6880},
						Fin:    Coordenada{Latitud: 37.2230, Longitud: -3.6880},
					}}},
					"Destino": {Tramos: []Tramo{{
						Inicio: Coordenada{Latitud: 37.2230, Longitud: -3.6880},
						Fin:    Coordenada{Latitud: 37.2240, Longitud: -3.6880},
					}}},
				},
			},
			origen:        "Origen",
			destino:       "Destino",
			errorEsperado: ErrTramoLongitudCero,
		},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, func(t *testing.T) {
			_, err := CalcularRuta(tc.callejero, tc.origen, tc.destino)
			if !errors.Is(err, tc.errorEsperado) {
				t.Errorf("error incorrecto: obtenido %v, esperado %v", err, tc.errorEsperado)
			}
		})
	}
}