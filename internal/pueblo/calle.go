package pueblo

type TipoCalle string

const (
	Ancha    TipoCalle = "ancha"
	Estrecha TipoCalle = "estrecha"
)

type Calle struct {
	Tipo             TipoCalle
	NumeroArboles    uint
	Inicio           Coordenada
	Fin              Coordenada
	PlantasEdificios []int
}
