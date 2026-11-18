package pueblo

import (
	"math"
	"testing"
)

func TestDistanciaEntreCoordenadas(t *testing.T) {
	const toleranciaMetros = 0.01

	inicioCalle := Coordenada{Latitud: 37.1760, Longitud: -3.5990}
	finCalle := Coordenada{Latitud: 37.1775, Longitud: -3.5975}
	origenEcuador := Coordenada{Latitud: 0.0, Longitud: 0.0}
	unGradoSur := Coordenada{Latitud: -1.0, Longitud: 0.0}

	casos := []struct {
		nombre            string
		desde             Coordenada
		hasta             Coordenada
		distanciaEsperada float64
	}{
		{
			nombre:            "la distancia de un punto a si mismo es cero",
			desde:             inicioCalle,
			hasta:             inicioCalle,
			distanciaEsperada: 0.0,
		},
		{
			nombre:            "calcula la longitud de un tramo urbano en metros",
			desde:             inicioCalle,
			hasta:             finCalle,
			distanciaEsperada: 213.26,
		},
		{
			nombre:            "la distancia es simetrica en sentido inverso",
			desde:             finCalle,
			hasta:             inicioCalle,
			distanciaEsperada: 213.26,
		},
		{
			nombre:            "calcula correctamente distancias hacia el hemisferio sur",
			desde:             origenEcuador,
			hasta:             unGradoSur,
			distanciaEsperada: 111194.93,
		},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, func(t *testing.T) {
			obtenida := tc.desde.DistanciaA(tc.hasta)

			if math.Abs(obtenida-tc.distanciaEsperada) > toleranciaMetros {
				t.Errorf("distancia incorrecta: obtenida %.2f m, esperada %.2f m", obtenida, tc.distanciaEsperada)
			}
		})
	}
}