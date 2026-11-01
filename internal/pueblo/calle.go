package pueblo

type Calle struct {
	Tipo      string
	Arboles   int
	Edificios map[Coordenada]Edificio
}
