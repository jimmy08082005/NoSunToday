package pueblo

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
)

const distanciaMaximaArbolMetros = 15.0

type respuestaOverpass struct {
	Elements []elementoOSM `json:"elements"`
}

type elementoOSM struct {
	Type  string            `json:"type"`
	ID    int64             `json:"id"`
	Lat   float64           `json:"lat,omitempty"`
	Lon   float64           `json:"lon,omitempty"`
	Nodes []int64           `json:"nodes,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
}

func CargarCalles(r io.Reader) (Callejero, error) {
	var datos respuestaOverpass
	if err := json.NewDecoder(r).Decode(&datos); err != nil {
		return Callejero{}, fmt.Errorf("error decodificando JSON de OpenStreetMap: %w", err)
	}

	nodos := make(map[int64]Coordenada)
	var arboles []Coordenada
	var vias []elementoOSM
	usoNodos := make(map[int64]int)

	for _, el := range datos.Elements {
		switch el.Type {
		case "node":
			coord := Coordenada{Latitud: el.Lat, Longitud: el.Lon}
			nodos[el.ID] = coord
			if el.Tags["natural"] == "tree" {
				arboles = append(arboles, coord)
			}
		case "way":
			if esViaPeatonalValida(el.Tags) {
				vias = append(vias, el)
				for _, nodeID := range el.Nodes {
					usoNodos[nodeID]++
				}
			}
		}
	}

	calles := make(map[NombreCalle]Calle)
	for _, via := range vias {
		if len(via.Nodes) < 2 {
			continue
		}

		nombre := NombreCalle(via.Tags["name"])
		inicioIdx := 0

		for i := 1; i < len(via.Nodes); i++ {
			nodeID := via.Nodes[i]
			esInterseccion := usoNodos[nodeID] > 1
			esFinal := i == len(via.Nodes)-1

			if esInterseccion || esFinal {
				coordInicio, okInicio := nodos[via.Nodes[inicioIdx]]
				coordFin, okFin := nodos[nodeID]

				if okInicio && okFin && coordInicio != coordFin {
					calle := calles[nombre]
					calle.Tramos = append(calle.Tramos, Tramo{
						Inicio: coordInicio,
						Fin:    coordFin,
					})
					calles[nombre] = calle
				}
				inicioIdx = i
			}
		}
	}

	asignarArbolesATramos(calles, arboles)

	return Callejero{Calles: calles}, nil
}

func esViaPeatonalValida(tags map[string]string) bool {
	highway, tieneTipo := tags["highway"]
	nombre := tags["name"]
	if !tieneTipo || nombre == "" {
		return false
	}
	switch highway {
	case "motorway", "motorway_link", "trunk", "trunk_link":
		return false
	default:
		return true
	}
}

func asignarArbolesATramos(calles map[NombreCalle]Calle, arboles []Coordenada) {
	for _, arbol := range arboles {
		var mejorCalle NombreCalle
		mejorIdx := -1
		menorDistancia := distanciaMaximaArbolMetros

		for nombre, calle := range calles {
			for i, tramo := range calle.Tramos {
				dist := distanciaPuntoASegmento(arbol, tramo.Inicio, tramo.Fin)
				if dist <= menorDistancia {
					menorDistancia = dist
					mejorCalle = nombre
					mejorIdx = i
				}
			}
		}

		if mejorIdx != -1 {
			calles[mejorCalle].Tramos[mejorIdx].NumeroArboles++
		}
	}
}

func distanciaPuntoASegmento(p, a, b Coordenada) float64 {
	cosLat := math.Cos(a.Latitud * math.Pi / 180.0)
	bx := (b.Longitud - a.Longitud) * cosLat
	by := b.Latitud - a.Latitud
	px := (p.Longitud - a.Longitud) * cosLat
	py := p.Latitud - a.Latitud

	largoCuadrado := bx*bx + by*by
	if largoCuadrado == 0 {
		return p.DistanciaA(a)
	}

	t := (px*bx + py*by) / largoCuadrado
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	masCercano := Coordenada{
		Latitud:  a.Latitud + t*(b.Latitud-a.Latitud),
		Longitud: a.Longitud + t*(b.Longitud-a.Longitud),
	}

	return p.DistanciaA(masCercano)
}