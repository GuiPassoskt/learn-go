package types

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

func (c Carro) buzinar() {
	print(c, "buzinou")
}
