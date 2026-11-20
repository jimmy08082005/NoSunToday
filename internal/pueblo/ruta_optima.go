package pueblo

func generarGrafo(callejero Callejero) map[Coordenada]map[NombreCalle][]Tramo {
	grafo := make(map[Coordenada]map[NombreCalle][]Tramo)

	for nombre, calle := range callejero.Calles {
		for _, tramo := range calle.Tramos {
			if tramo.Inicio == tramo.Fin {
				continue
			}

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