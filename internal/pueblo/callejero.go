package pueblo

type NombreCalle string

type Callejero struct {
	Calles map[NombreCalle]Calle
}
