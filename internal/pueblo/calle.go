package pueblo

type TipoCalle string

const (
	Ancha    TipoCalle = "ancha"
	Estrecha TipoCalle = "estrecha"
)

type Edificio struct {
	Plantas   int
	Ubicacion Coordenada
}

type Calle struct {
	Tipo          TipoCalle
	NumeroArboles uint
	Inicio        Coordenada
	Fin           Coordenada
	Edificios     []Edificio
}
